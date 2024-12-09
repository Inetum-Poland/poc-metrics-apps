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
	"inetum.com/metrics-go-app/internal/config"
)

var (
	Tracer        trace.Tracer
	TraceProvider *sdktrace.TracerProvider

	Meter         metric.Meter
	MeterProvider *sdkmetric.MeterProvider

	Logger         *slog.Logger
	LoggerProvider *sdklog.LoggerProvider

	serviceName  string
	collectorURL string

	ApiCounter  metric.Int64Counter
	DbCounter   metric.Int64Counter
	FailCounter metric.Int64Counter
	ApiReqCount metric.Int64Counter
	ApiReqTime  metric.Float64Histogram
)

func setupMetics() {
	ApiCounter, _ = Meter.Int64Counter(
		"api.counter",
		metric.WithDescription("Number of API calls."),
		metric.WithUnit("{call}"),
	)

	DbCounter, _ = Meter.Int64Counter(
		"db.counter",
		metric.WithDescription("Number of DB calls."),
		metric.WithUnit("{call}"),
	)

	ApiReqCount, _ = Meter.Int64Counter(
		"api.req.count",
		metric.WithDescription("Number of requests."),
	)

	ApiReqTime, _ = Meter.Float64Histogram(
		"api.req.time",
		metric.WithDescription("Request time."),
		metric.WithExplicitBucketBoundaries(0.001, 0.005, 0.010, 0.050, 0.100, 0.500, 1.000, 5.000, 10.000),
	)
}

func Setup() {
	collectorURL = fmt.Sprintf("%s:%d", config.C.Otel.Host, config.C.Otel.Port)
	serviceName = config.C.WebApp.Name

	ctx := context.Background()
	TraceProvider, Tracer = SetupTracer(ctx)
	MeterProvider, Meter = SetupMeter(ctx)
	LoggerProvider, Logger = SetupLogger(ctx)

	setupMetics()
}
