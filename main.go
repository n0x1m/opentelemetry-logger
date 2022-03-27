package main

import (
	"context"

	"github.com/n0x1m/opentelemetry-logger/logger"
	"github.com/n0x1m/opentelemetry-logger/tracer"
	"go.opentelemetry.io/otel/trace"
	"go.uber.org/zap/zapcore"
)

func main() {
	log, _ := logger.New(zapcore.InfoLevel, true)
	ctx := logger.NewContext(context.Background(), log)

	tp := tracer.NewTraceProvider()
	tc := tp.Tracer("opentelemetry-logger/main")
	ctx = tracer.NewContext(ctx, tc)

	logger.FromContext(ctx).Info("start test without traces")

	// pass context
	run(ctx)
	run2(ctx)

	logger.FromContext(ctx).Info("exit test without traces")
}

func run(ctx context.Context) {
	var span trace.Span
	ctx, span = tracer.FromContext(ctx).Start(ctx, "run")
	defer span.End()

	logger.FromContext(ctx).Info("span test 1 with trace")
	logger.FromContext(ctx).Info("span test 1 repeats trace and span id")
}

func run2(ctx context.Context) {
	var span trace.Span
	ctx, span = tracer.FromContext(ctx).Start(ctx, "run2")
	defer span.End()

	logger.FromContext(ctx).Info("span test 2 with new trace")
	run3(ctx)
}

func run3(ctx context.Context) {
	var span trace.Span
	ctx, span = tracer.FromContext(ctx).Start(ctx, "run3")
	defer span.End()

	logger.FromContext(ctx).Info("span test 3 repeates trace id from 2 with new span id")
}
