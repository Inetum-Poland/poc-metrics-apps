package otel

import (
	"context"
	"fmt"
	__log "log"
	"log/slog"

	otelslog "go.opentelemetry.io/contrib/bridges/otelslog"
	"go.opentelemetry.io/otel/exporters/otlp/otlplog/otlploggrpc"
	"go.opentelemetry.io/otel/exporters/otlp/otlplog/otlploghttp"
	sdklog "go.opentelemetry.io/otel/sdk/log"
	"go.opentelemetry.io/otel/sdk/resource"
	semconv "go.opentelemetry.io/otel/semconv/v1.26.0"
)

func newHttpLogExporter(ctx context.Context) (sdklog.Exporter, error) {
	return otlploghttp.New(ctx,
		otlploghttp.WithEndpoint(fmt.Sprintf("%s:%d", collectorURL, 4318)),
		otlploghttp.WithInsecure(),
	)
}

func newGrpcLogExporter(ctx context.Context) (sdklog.Exporter, error) {
	return otlploggrpc.New(ctx,
		otlploggrpc.WithEndpoint(fmt.Sprintf("%s:%d", collectorURL, 4317)),
		otlploggrpc.WithInsecure(),
	)
}

func newLogProvider(exp sdklog.Exporter) *sdklog.LoggerProvider {
	r, err := resource.Merge(
		resource.Default(),
		resource.NewWithAttributes(
			semconv.SchemaURL,
			semconv.ServiceName(serviceName),
			semconv.DeploymentEnvironment("local"),
		),
	)

	if err != nil {
		panic(err)
	}

	return sdklog.NewLoggerProvider(
		sdklog.WithResource(r),
		sdklog.WithProcessor(
			sdklog.NewBatchProcessor(exp),
		),
	)
}

func SetupLogger(ctx context.Context) (*sdklog.LoggerProvider, *slog.Logger) {
	exp, err := newHttpLogExporter(ctx)
	if err != nil {
		__log.Fatalf("failed to initialize exporter: %v", err)
	}

	logProvider := newLogProvider(exp)
	// logger := logProvider.Logger("pl.inetum.com/go-otel-log")
	logger := otelslog.NewLogger(
		serviceName,
		otelslog.WithLoggerProvider(logProvider),
	)

	return logProvider, logger
}
