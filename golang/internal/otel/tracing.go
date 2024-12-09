package otel

import (
	"context"
	"log"

	"go.opentelemetry.io/otel"
	"go.opentelemetry.io/otel/exporters/otlp/otlptrace/otlptracegrpc"
	"go.opentelemetry.io/otel/sdk/resource"
	sdktrace "go.opentelemetry.io/otel/sdk/trace"
	semconv "go.opentelemetry.io/otel/semconv/v1.26.0"
	trace "go.opentelemetry.io/otel/trace"
)

// func newHttpTraceExporter(ctx context.Context) (sdktrace.SpanExporter, error) {
// 	return otlptracehttp.New(ctx,
// 		otlptracehttp.WithEndpoint(fmt.Sprintf("%s:%d", collectorURL, 4318)),
// 		otlptracehttp.WithInsecure(),
// 	)
// }

func newGrpcTraceExporter(ctx context.Context) (sdktrace.SpanExporter, error) {
	return otlptracegrpc.New(ctx,
		otlptracegrpc.WithEndpoint(collectorURL),
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

func SetupTracer(ctx context.Context) (*sdktrace.TracerProvider, trace.Tracer) {
	exp, err := newGrpcTraceExporter(ctx)
	// exp, err := newHttpTraceExporter(ctx)
	if err != nil {
		log.Fatalf("failed to initialize exporter: %v", err)
	}

	traceProvider := newTraceProvider(exp)
	otel.SetTracerProvider(traceProvider)
	tracer := traceProvider.Tracer("pl.inetum.com/go-otel-tracing",
		trace.WithInstrumentationVersion("0.0.1"))

	return traceProvider, tracer
}
