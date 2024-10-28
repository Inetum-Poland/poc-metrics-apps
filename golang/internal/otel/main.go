package otel

import (
	"context"
	"fmt"

	sdktrace "go.opentelemetry.io/otel/sdk/trace"
	"go.opentelemetry.io/otel/trace"
)

var Tracer trace.Tracer
var TraceProvider *sdktrace.TracerProvider

var serviceName = "AppGolang"
var collectorURL = "opentelemetry"

var otlpHeaders = map[string]string{
	"X-Otel-Trace-Exporter":          "otlp",
	"X-Otel-Trace-Exporter-Endpoint": fmt.Sprintf("%s:%d", collectorURL, 4317),
}

// Singleton; Runs only on the first import.
func init() {
	TraceProvider, Tracer = SetupTracer(context.Background())
}
