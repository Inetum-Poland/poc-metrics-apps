package main

import (
	"context"
	"fmt"
	"log"

	"go.opentelemetry.io/otel"
	"go.opentelemetry.io/otel/exporters/otlp/otlptrace/otlptracegrpc"
	"go.opentelemetry.io/otel/exporters/otlp/otlptrace/otlptracehttp"
	"go.opentelemetry.io/otel/sdk/resource"
	sdktrace "go.opentelemetry.io/otel/sdk/trace"
	semconv "go.opentelemetry.io/otel/semconv/v1.26.0"
	"go.opentelemetry.io/otel/trace"
)

var serviceName = "AppGolang"
var collectorURL = "opentelemetry"

var otlpHeaders = map[string]string{
	"X-Otel-Trace-Exporter":          "otlp",
	"X-Otel-Trace-Exporter-Endpoint": fmt.Sprintf("%s:%d", collectorURL, 4317),
}

var tracer trace.Tracer

func newHttpTraceExporter(ctx context.Context) (sdktrace.SpanExporter, error) {
	return otlptracehttp.New(ctx,
		otlptracehttp.WithEndpoint(fmt.Sprintf("%s:%d", collectorURL, 4318)),
		otlptracehttp.WithHeaders(otlpHeaders),
		otlptracehttp.WithInsecure(),
	)
}

func newGrpcTraceExporter(ctx context.Context) (sdktrace.SpanExporter, error) {
	return otlptracegrpc.New(ctx,
		otlptracegrpc.WithEndpoint(fmt.Sprintf("%s:%d", collectorURL, 4317)),
		otlptracegrpc.WithHeaders(otlpHeaders),
		otlptracegrpc.WithInsecure(),
	)
}

func newTraceProvider(exp sdktrace.SpanExporter) *sdktrace.TracerProvider {
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

	return sdktrace.NewTracerProvider(
		sdktrace.WithBatcher(exp),
		sdktrace.WithResource(r),
	)
}

func setupTracer(ctx context.Context) *sdktrace.TracerProvider {
	// exp, err := newGrpcTraceExporter(ctx)
	exp, err := newHttpTraceExporter(ctx)
	if err != nil {
		log.Fatalf("failed to initialize exporter: %v", err)
	}

	traceProvider := newTraceProvider(exp)
	otel.SetTracerProvider(traceProvider)
	tracer = traceProvider.Tracer("pl.inetum.com/go-otel-tracing")

	return traceProvider
}
