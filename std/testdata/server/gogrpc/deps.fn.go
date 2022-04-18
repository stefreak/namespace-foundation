// This file was automatically generated by Foundation.
// DO NOT EDIT. To update, re-run `fn generate`.

package main

import (
	"context"
	"namespacelabs.dev/foundation/std/go/core"
	"namespacelabs.dev/foundation/std/go/grpc/metrics"
	"namespacelabs.dev/foundation/std/go/server"
	"namespacelabs.dev/foundation/std/grpc/deadlines"
	"namespacelabs.dev/foundation/std/grpc/logging"
	"namespacelabs.dev/foundation/std/monitoring/tracing"
	"namespacelabs.dev/foundation/std/testdata/service/post"
	"namespacelabs.dev/foundation/universe/go/panicparse"
)

func RegisterInitializers(di *core.DependencyGraph) {
	di.AddInitializers(metrics.Initializers__so2f3v...)
	di.AddInitializers(tracing.Initializers__70o2mm...)
	di.AddInitializers(deadlines.Initializers__vbko45...)
	di.AddInitializers(logging.Initializers__16bc0q...)
	di.AddInitializers(panicparse.Initializers__99b5nh...)
}

func WireServices(ctx context.Context, srv server.Server, depgraph core.Dependencies) []error {
	var errs []error

	if err := depgraph.Instantiate(ctx, post.Provider__j7h7h5, func(ctx context.Context, v interface{}) error {
		post.WireService(ctx, srv.Scope(post.Package__j7h7h5), v.(post.ServiceDeps))
		return nil
	}); err != nil {
		errs = append(errs, err)
	}

	srv.InternalRegisterGrpcGateway(post.RegisterPostServiceHandler)

	return errs
}
