// Copyright 2022 Namespace Labs Inc; All rights reserved.
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.

package cluster

import (
	"context"
	"encoding/json"
	"fmt"
	"net"
	"net/url"
	"regexp"
	"strconv"
	"strings"

	"github.com/containerd/containerd"
	"github.com/containerd/containerd/namespaces"
	"github.com/containerd/nerdctl/pkg/idutil/containerwalker"
	"github.com/containerd/nerdctl/pkg/labels"
	"github.com/containerd/nerdctl/pkg/portutil"
	"github.com/docker/docker/client"
	"github.com/spf13/cobra"
	"golang.org/x/exp/slices"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/credentials/insecure"
	"namespacelabs.dev/foundation/framework/rpcerrors"
	"namespacelabs.dev/foundation/internal/cli/fncobra"
	"namespacelabs.dev/foundation/internal/console"
	"namespacelabs.dev/foundation/internal/fnapi"
	"namespacelabs.dev/foundation/internal/fnerrors"
	"namespacelabs.dev/foundation/internal/providers/nscloud/api"
)

func NewExposeCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "expose [cluster-id]",
		Short: "Opens a public ingress to the specified exported port.",
		Args:  cobra.MaximumNArgs(1),
	}

	source := cmd.Flags().String("source", "docker", "Where to lookup the container.")
	prefix := cmd.Flags().String("prefix", "", "If specified, prefixes the allocated URL.")
	containerName := cmd.Flags().String("container", "", "Which container to export.")
	containerPorts := cmd.Flags().IntSlice("container_port", nil, "If specified, only exposes the specified ports.")
	output := cmd.Flags().StringP("output", "o", "plain", "One of plain or json.")

	cmd.RunE = fncobra.RunE(func(ctx context.Context, args []string) error {
		if *containerName == "" {
			return fnerrors.New("--container is required")
		}

		cluster, _, err := selectRunningCluster(ctx, args)
		if err != nil {
			return err
		}

		if cluster == nil {
			return nil
		}

		ports, err := selectPorts(ctx, cluster, *source, *containerName)
		if err != nil {
			return err
		}

		if len(*containerPorts) > 0 {
			remapped := portMap{}

			for _, port := range *containerPorts {
				if m, has := ports[port]; !has {
					return fnerrors.New("port %d not exported by container", port)
				} else {
					remapped[port] = m
				}
			}

			ports = remapped
		}

		var exps []exported
		for containerPort, port := range ports {
			p := *prefix
			if p == "" {
				p = port.SuggestedPrefix
			}

			resp, err := api.RegisterDefaultIngress(ctx, api.Endpoint, api.RegisterDefaultIngressRequest{
				ClusterId: cluster.ClusterId,
				Prefix:    p,
				BackendEndpoint: &api.IngressBackendEndpoint{
					Port: port.ExportedPort,
				},
			})
			if err != nil {
				return err
			}

			exps = append(exps, exported{
				ContainerID:   port.ContainerID,
				ContainerName: port.ContainerName,
				ContainerPort: int32(containerPort),
				URL:           "https://" + resp.Fqdn,
			})

			if *output == "plain" {
				fmt.Fprintf(console.Stdout(ctx), "Exported port %d from %s (%s):\n  https://%s\n\n",
					containerPort, port.ContainerName, substr(port.ContainerID), resp.Fqdn)
			}
		}

		if *output == "json" {
			slices.SortFunc(exps, func(a, b exported) bool {
				return a.ContainerPort < b.ContainerPort
			})

			return json.NewEncoder(console.Stdout(ctx)).Encode(exps)
		}

		return nil
	})

	return cmd
}

type exported struct {
	ContainerID   string `json:"container_id"`
	ContainerName string `json:"container_name"`
	ContainerPort int32  `json:"container_port"`
	URL           string `json:"url"`
}

type containerPort struct {
	ContainerID     string
	ContainerName   string
	SuggestedPrefix string
	ExportedPort    int32
}

type portMap map[int]containerPort

func selectPorts(ctx context.Context, cluster *api.KubernetesCluster, source, containerName string) (portMap, error) {
	switch source {
	case "docker":
		return selectDockerPorts(ctx, cluster, containerName)

	case "containerd":
		return selectContainerdPorts(ctx, cluster, containerName)
	}

	return nil, fnerrors.New("unsupported source %q", source)
}

func selectDockerPorts(ctx context.Context, cluster *api.KubernetesCluster, containerName string) (portMap, error) {
	// We must fetch a token with our parent context, so we get a task sink etc.
	token, err := fnapi.FetchTenantToken(ctx)
	if err != nil {
		return nil, err
	}

	docker, err := client.NewClientWithOpts(client.WithDialContext(func(ctx context.Context, network, addr string) (net.Conn, error) {
		return api.DialPortWithToken(ctx, token, cluster, 2375)
	}))
	if err != nil {
		return nil, err
	}

	defer docker.Close()

	data, err := docker.ContainerInspect(ctx, containerName)
	if err != nil {
		return nil, err
	}

	internalName, suggestedPrefix := parseContainerName(data.ID, data.Name)

	exported := portMap{}
	for port, mapping := range data.NetworkSettings.Ports {
		if port.Proto() == "tcp" && len(mapping) > 0 {
			for _, m := range mapping {
				if m.HostIP == "0.0.0.0" {
					parsedPort, err := strconv.ParseInt(m.HostPort, 10, 32)
					if err != nil {
						return nil, err
					}

					exported[port.Int()] = containerPort{
						ContainerID:     data.ID,
						ContainerName:   internalName,
						SuggestedPrefix: suggestedPrefix,
						ExportedPort:    int32(parsedPort),
					}
				} else {
					fmt.Fprintf(console.Warnings(ctx), "Skipping %d/%s exported to %s (unsupported)\n", port.Int(), port.Proto(), m.HostIP)
				}
			}
		} else {
			fmt.Fprintf(console.Warnings(ctx), "Skipping %d/%s (unsupported)\n", port.Int(), port.Proto())
		}
	}

	return exported, nil
}

func parseContainerName(id, name string) (string, string) {
	internalName := strings.TrimPrefix(name, "/")             // docker returns names prefixed by /
	mangledName := strings.ReplaceAll(internalName, "_", "-") // docker generated names have underscores.

	suggestedPrefix := computeSuggestedPrefix(id, mangledName)

	return internalName, suggestedPrefix
}

func withContainerd(ctx context.Context, cluster *api.KubernetesCluster, callback func(context.Context, *containerd.Client) error) error {
	// We must fetch a token with our parent context, so we get a task sink etc.
	token, err := fnapi.FetchTenantToken(ctx)
	if err != nil {
		return err
	}

	conn, err := grpc.DialContext(ctx, fmt.Sprintf("%s-containerd", cluster.ClusterId),
		grpc.WithBlock(),
		grpc.WithContextDialer(func(ctx context.Context, s string) (net.Conn, error) {
			vars := url.Values{}
			vars.Set("name", "containerd-socket")
			return api.DialHostedServiceWithToken(ctx, token, cluster, "unixsocket", vars)
		}),
		grpc.WithTransportCredentials(insecure.NewCredentials()),
	)
	if err != nil {
		return err
	}

	defer conn.Close()

	ctr, err := containerd.NewWithConn(conn)
	if err != nil {
		return err
	}

	defer ctr.Conn()

	ctx = namespaces.WithNamespace(ctx, "default")

	return callback(ctx, ctr)
}

func selectContainerdPorts(ctx context.Context, cluster *api.KubernetesCluster, containerName string) (portMap, error) {
	exported := portMap{}

	if err := withContainerd(ctx, cluster, func(ctx context.Context, ctr *containerd.Client) error {
		walker := &containerwalker.ContainerWalker{
			Client: ctr,
			OnFound: func(ctx context.Context, found containerwalker.Found) error {
				if found.MatchCount > 1 {
					return rpcerrors.Errorf(codes.InvalidArgument, "container name matches multiple containers")
				}

				l, err := found.Container.Labels(ctx)
				if err != nil {
					return err
				}

				ports, err := portutil.ParsePortsLabel(l)
				if err != nil {
					return err
				}

				internalName, suggestedPrefix := parseContainerName(found.Container.ID(), l[labels.Name])

				for _, p := range ports {
					if p.Protocol == "tcp" && p.HostIP == "0.0.0.0" {
						exported[int(p.ContainerPort)] = containerPort{
							ContainerID:     found.Container.ID(),
							ContainerName:   internalName,
							SuggestedPrefix: suggestedPrefix,
							ExportedPort:    p.HostPort,
						}
					} else {
						fmt.Fprintf(console.Warnings(ctx), "Skipping %d/%s exported to %s (unsupported)\n", p.ContainerPort, p.Protocol, p.HostIP)
					}
				}

				return nil
			},
		}

		n, err := walker.Walk(ctx, containerName)
		if err != nil {
			return err
		} else if n == 0 {
			return rpcerrors.Errorf(codes.NotFound, "no such container %q", containerName)
		}

		return nil
	}); err != nil {
		return nil, err
	}

	return exported, nil
}

var simpleLabelRe = regexp.MustCompile("^[a-zA-Z0-9][a-zA-Z0-9-]*$")

func computeSuggestedPrefix(id, name string) string {
	if len(name) < 24 && simpleLabelRe.MatchString(name) {
		return name
	}

	return substr(id)
}

func substr(id string) string {
	if len(id) > 8 {
		return id[8:]
	}

	return id
}
