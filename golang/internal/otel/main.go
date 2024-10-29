package otel

import (
	"context"
	"fmt"
	"log/slog"

	metric "go.opentelemetry.io/otel/metric"
	sdklog "go.opentelemetry.io/otel/sdk/log"
	sdkmetric "go.opentelemetry.io/otel/sdk/metric"
	sdktrace "go.opentelemetry.io/otel/sdk/trace"
	trace "go.opentelemetry.io/otel/trace"
)

var Tracer trace.Tracer
var TraceProvider *sdktrace.TracerProvider

var Meter metric.Meter
var MeterProvider *sdkmetric.MeterProvider

// var SLogger *slog.Logger
var Logger *slog.Logger
var LoggerProvider *sdklog.LoggerProvider

var serviceName = "AppGolang"
var collectorURL = "opentelemetry"

var otlpHttpConnection = fmt.Sprintf("%s:%d", collectorURL, 4318)
var otlpGrpcConnection = fmt.Sprintf("%s:%d", collectorURL, 4317)

// Singleton; Runs only on the first import.
func init() {
	TraceProvider, Tracer = SetupTracer(context.Background())
	MeterProvider, Meter = SetupMeter(context.Background())
	LoggerProvider, Logger = SetupLogger(context.Background())
}
