// Copyright 2022 Namespace Labs Inc; All rights reserved.
// Licensed under the EARLY ACCESS SOFTWARE LICENSE AGREEMENT
// available at http://github.com/namespacelabs/foundation

package orchestration

import (
	"context"

	"namespacelabs.dev/foundation/internal/engine/ops"
	"namespacelabs.dev/foundation/internal/fnerrors"
	"namespacelabs.dev/foundation/provision/deploy"
	"namespacelabs.dev/foundation/runtime"
	"namespacelabs.dev/foundation/schema"
	orchpb "namespacelabs.dev/foundation/schema/orchestration"
	"namespacelabs.dev/foundation/std/planning"
	"namespacelabs.dev/foundation/workspace/compute"
)

func Deploy(ctx context.Context, env planning.Context, cluster runtime.Cluster, p *ops.Plan, plan *schema.DeployPlan, wait, outputProgress bool) error {
	if UseOrchestrator {
		cli, err := compute.GetValue(ctx, ConnectToClient(env, cluster))
		if err != nil {
			return err
		}

		id, err := CallDeploy(ctx, env, cli, plan)
		if err != nil {
			return err
		}

		if wait {
			if outputProgress {
				return deploy.RenderAndWait(ctx, env, cluster, func(ch chan *orchpb.Event) error {
					return WireDeploymentStatus(ctx, cli, id, ch)
				})
			} else {
				return WireDeploymentStatus(ctx, cli, id, nil)
			}
		}
	} else {
		waiters, err := ops.Execute(ctx, runtime.TaskServerDeploy, env, p)
		if err != nil {
			return fnerrors.New("failed to deploy: %w", err)
		}

		if wait {
			if outputProgress {
				return deploy.Wait(ctx, env, cluster, waiters)
			} else {
				return ops.WaitMultiple(ctx, waiters, nil)
			}
		}
	}

	return nil
}
