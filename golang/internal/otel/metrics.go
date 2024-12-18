package otel

import (
	"context"
	"log"
	"time"

	"go.opentelemetry.io/otel"
	"go.opentelemetry.io/otel/exporters/otlp/otlpmetric/otlpmetricgrpc"
	metric "go.opentelemetry.io/otel/metric"
	sdkmetric "go.opentelemetry.io/otel/sdk/metric"
	"go.opentelemetry.io/otel/sdk/resource"
	semconv "go.opentelemetry.io/otel/semconv/v1.26.0"
)

// func newHttpMerticExporter(ctx context.Context) (sdkmetric.Exporter, error) {
// 	return otlpmetrichttp.New(ctx,
// 		otlpmetrichttp.WithEndpoint(fmt.Sprintf("%s:%d", collectorURL, 4318)),
// 		otlpmetrichttp.WithInsecure(),
// 	)
// }

func newGrpcMetricExporter(ctx context.Context) (sdkmetric.Exporter, error) {
	return otlpmetricgrpc.New(ctx,
		otlpmetricgrpc.WithEndpoint(collectorURL),
		otlpmetricgrpc.WithInsecure(),
	)
}

func newMeterProvider(exp sdkmetric.Exporter) *sdkmetric.MeterProvider {
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

	return sdkmetric.NewMeterProvider(
		sdkmetric.WithResource(r),
		sdkmetric.WithReader(
			sdkmetric.NewPeriodicReader(
				exp,
				sdkmetric.WithInterval(10*time.Second),
			),
		),
	)
}

func SetupMeter(ctx context.Context) (*sdkmetric.MeterProvider, metric.Meter) {
	exp, err := newGrpcMetricExporter(ctx)
	// exp, err := newHttpMerticExporter(ctx)
	if err != nil {
		log.Fatalf("failed to initialize exporter: %v", err)
	}

	meterProvider := newMeterProvider(exp)
	otel.SetMeterProvider(meterProvider)
	meter := meterProvider.Meter("pl.inetum.com/go-otel-metrics",
		metric.WithInstrumentationVersion("0.0.1"))

	return meterProvider, meter
}
