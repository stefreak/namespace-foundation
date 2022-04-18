// This file was automatically generated by Foundation.
// DO NOT EDIT. To update, re-run `fn generate`.

package main

import (
	"context"
	"namespacelabs.dev/foundation/std/go/core"
	"namespacelabs.dev/foundation/std/go/grpc/metrics"
	"namespacelabs.dev/foundation/std/go/server"
	"namespacelabs.dev/foundation/std/monitoring/tracing"
	"namespacelabs.dev/foundation/std/testdata/service/multicounter"
)

func RegisterInitializers(di *core.DependencyGraph) {
	di.AddInitializers(metrics.Initializers__so2f3v...)
	di.AddInitializers(tracing.Initializers__70o2mm...)
}

func WireServices(ctx context.Context, srv server.Server, depgraph core.Dependencies) []error {
	var errs []error

	if err := depgraph.Instantiate(ctx, multicounter.Provider__n4bhfe, func(ctx context.Context, v interface{}) error {
		multicounter.WireService(ctx, srv.Scope(multicounter.Package__n4bhfe), v.(multicounter.ServiceDeps))
		return nil
	}); err != nil {
		errs = append(errs, err)
	}

	return errs
}
