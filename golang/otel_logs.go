package main

import (
	"context"

	"go.opentelemetry.io/otel/exporters/otlp/otlplog/otlploggrpc"
	"go.opentelemetry.io/otel/exporters/otlp/otlplog/otlploghttp"
	"go.opentelemetry.io/otel/sdk/log"
)

func newHttpLogExporter(ctx context.Context) (log.Exporter, error) {
	return otlploghttp.New(ctx)
}

func newGrpcLogExporter(ctx context.Context) (log.Exporter, error) {
	return otlploggrpc.New(ctx)
}
