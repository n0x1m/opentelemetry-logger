package tracer

import (
	"context"
	"log"

	"go.opentelemetry.io/otel"
	"go.opentelemetry.io/otel/exporters/stdout/stdouttrace"
	sdktrace "go.opentelemetry.io/otel/sdk/trace"
	"go.opentelemetry.io/otel/trace"
)

func NewTraceProvider() *sdktrace.TracerProvider {
	var err error
	exp, err := stdouttrace.New(stdouttrace.WithPrettyPrint())
	if err != nil {
		log.Panicf("failed to initialize stdouttrace exporter %v\n", err)
		return nil
	}
	bsp := sdktrace.NewBatchSpanProcessor(exp)
	tp := sdktrace.NewTracerProvider(
		sdktrace.WithSampler(sdktrace.AlwaysSample()),
		sdktrace.WithSpanProcessor(bsp),
	)
	otel.SetTracerProvider(tp)

	return tp
}

type ctxKey struct{}

func FromContext(ctx context.Context) trace.Tracer {
	t, _ := ctx.Value(ctxKey{}).(trace.Tracer)
	return t
}

func NewContext(parent context.Context, t trace.Tracer) context.Context {
	return context.WithValue(parent, ctxKey{}, t)
}
