// Copyright 2022 Namespace Labs Inc; All rights reserved.
// Licensed under the EARLY ACCESS SOFTWARE LICENSE AGREEMENT
// available at http://github.com/namespacelabs/foundation

package runtime

import (
	"context"
	"fmt"

	"namespacelabs.dev/foundation/internal/console"
	"namespacelabs.dev/foundation/internal/engine/ops"
	"namespacelabs.dev/foundation/internal/fnerrors"
	"namespacelabs.dev/foundation/schema"
	"namespacelabs.dev/foundation/std/planning"
)

func RunAttached(ctx context.Context, config planning.Configuration, cluster ClusterNamespace, spec DeployableSpec, io TerminalIO) error {
	spec.Attachable = io.TTY

	plan, err := cluster.Planner().PlanDeployment(ctx, DeploymentSpec{
		Specs: []DeployableSpec{spec},
	})
	if err != nil {
		return err
	}

	g := ops.NewPlan()
	if err := g.Add(plan.Definitions...); err != nil {
		return err
	}

	defer func() {
		if err := cluster.DeleteDeployment(ctx, spec); err != nil {
			fmt.Fprintf(console.Errors(ctx), "Deleting %s failed: %v\n", spec.Name, err)
		}
	}()

	// ResolveContainers will wait until the deployable is running, so we don't rely on the waiters returned by Execute.
	if _, err := ops.Execute(ctx, config, "deployable.run-attached", g, ClusterInjection.With(cluster.Cluster())); err != nil {
		return fnerrors.New("failed to deploy: %w", err)
	}

	containers, err := cluster.ResolveContainers(ctx, spec)
	if err != nil {
		return err
	}

	var mainContainers []*ContainerReference
	for _, container := range containers {
		if container.Kind == schema.ContainerKind_PRIMARY {
			mainContainers = append(mainContainers, container)
		}
	}

	if len(mainContainers) != 1 {
		return fnerrors.InternalError("expected a single container, saw %d", len(mainContainers))
	}

	return cluster.Cluster().AttachTerminal(ctx, mainContainers[0], io)
}
