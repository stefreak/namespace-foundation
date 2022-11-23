// Copyright 2022 Namespace Labs Inc; All rights reserved.
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.

package nscloud

import (
	"context"

	"namespacelabs.dev/foundation/internal/build/buildkit"
	"namespacelabs.dev/foundation/internal/fnerrors"
	"namespacelabs.dev/foundation/internal/providers/nscloud/api"
)

func EnsureBuildCluster(ctx context.Context) (*buildkit.Overrides, error) {
	cfg, err := api.CreateAndWaitCluster(ctx, "16x32", false, "build machine", []string{"BUILD_CLUSTER"})
	if err != nil {
		return nil, err
	}

	if cfg.BuildCluster != nil {
		return &buildkit.Overrides{
			HostedBuildCluster: &buildkit.HostedBuildCluster{
				ClusterId:  cfg.BuildCluster.Colocated.ClusterId,
				TargetPort: cfg.BuildCluster.Colocated.TargetPort,
			},
		}, nil
	} else {
		return nil, fnerrors.InternalError("expected build machine")
	}
}
