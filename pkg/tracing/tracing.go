// pkg/tracing/tracing.go
// Used by: All microservices (for request correlation)

package tracing

import (
	"context"

	"go.opentelemetry.io/otel"
	"go.opentelemetry.io/otel/sdk/trace"
)

func Init(service string) func(context.Context) {
	tp := trace.NewTracerProvider()
	otel.SetTracerProvider(tp)
	tracer := otel.Tracer(service)
	return func(ctx context.Context) {
		spanCtx, span := tracer.Start(ctx, service)
		defer span.End()
		_ = spanCtx
	}
}
