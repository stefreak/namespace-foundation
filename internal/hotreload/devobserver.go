// Copyright 2022 Namespace Labs Inc; All rights reserved.
// Licensed under the EARLY ACCESS SOFTWARE LICENSE AGREEMENT
// available at http://github.com/namespacelabs/foundation

package hotreload

import (
	"context"
	"fmt"
	"io"
	"strings"
	"sync"
	"time"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"namespacelabs.dev/foundation/internal/console"
	"namespacelabs.dev/foundation/internal/fnerrors/multierr"
	"namespacelabs.dev/foundation/internal/fnfs/workspace/wsremote"
	"namespacelabs.dev/foundation/internal/uniquestrings"
	"namespacelabs.dev/foundation/internal/wscontents"
	"namespacelabs.dev/foundation/provision"
	"namespacelabs.dev/foundation/runtime"
	"namespacelabs.dev/foundation/schema"
)

type FileSyncDevObserver struct {
	ctx          context.Context
	log          io.Writer
	server       *schema.Server
	cluster      runtime.ClusterNamespace
	fileSyncPort int32

	mu   sync.Mutex
	conn *grpc.ClientConn
	port io.Closer
}

func NewFileSyncDevObserver(ctx context.Context, cluster runtime.ClusterNamespace, srv provision.Server, fileSyncPort int32) *FileSyncDevObserver {
	return &FileSyncDevObserver{
		ctx:          ctx,
		log:          console.TypedOutput(ctx, "hot reload", console.CatOutputUs),
		server:       srv.Proto(),
		cluster:      cluster,
		fileSyncPort: fileSyncPort,
	}
}

func (do *FileSyncDevObserver) Close() error {
	do.mu.Lock()
	defer do.mu.Unlock()
	return do.cleanup()
}

func (do *FileSyncDevObserver) cleanup() error {
	var errs []error

	if do.conn != nil {
		if err := do.conn.Close(); err != nil {
			errs = append(errs, err)
		}
		do.conn = nil
	}

	if do.port != nil {
		if err := do.port.Close(); err != nil {
			errs = append(errs, err)
		}
		do.port = nil
	}

	return multierr.New(errs...)
}

func (do *FileSyncDevObserver) OnDeployment() {
	do.mu.Lock()
	defer do.mu.Unlock()

	err := do.cleanup()
	if err != nil {
		fmt.Fprintln(do.log, "failed to port forwarding cleanup", err)
	}

	// An endpoint is manufactored here, we don't actually export this in our metadata.
	do.port, err = do.cluster.ForwardPort(do.ctx, do.server, do.fileSyncPort, []string{"127.0.0.1"}, func(fp runtime.ForwardedPort) {
		do.onEndpoint(fp)
	})
	if err != nil {
		fmt.Fprintln(do.log, "failed to port forward filesync port", err)
	}
}

func (do *FileSyncDevObserver) onEndpoint(fpe runtime.ForwardedPort) {
	do.mu.Lock()
	defer do.mu.Unlock()

	if do.conn != nil {
		do.conn.Close()
		do.conn = nil
	}

	host := fmt.Sprintf("127.0.0.1:%d", fpe.LocalPort)

	conn, err := grpc.DialContext(do.ctx, host, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		fmt.Fprintln(do.log, "failed to connect to filesync", err)
	} else {
		do.conn = conn

		fmt.Fprintln(do.log, "Connected to FileSync (for hot reload).")
	}
}

func (do *FileSyncDevObserver) Deposit(ctx context.Context, s *wsremote.Signature, fe []*wscontents.FileEvent) error {
	do.mu.Lock()
	defer do.mu.Unlock()

	if do.conn == nil {
		return wsremote.ErrNotReady
	}

	var paths uniquestrings.List
	for _, r := range fe {
		paths.Add(r.Path)
	}

	fmt.Fprintf(do.log, "FileSync event: %s, paths: %s\n", s, strings.Join(paths.Strings(), ", "))

	newCtx, cancel := context.WithTimeout(ctx, 30*time.Second)
	defer cancel()

	if _, err := wsremote.NewFileSyncServiceClient(do.conn).Push(newCtx, &wsremote.PushRequest{
		Signature: s,
		FileEvent: fe,
	}); err != nil {
		return err
	}

	return nil
}
