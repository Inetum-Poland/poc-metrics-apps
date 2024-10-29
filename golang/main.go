package main

import (
	"context"
	"log/slog"

	api "inetum.com/metrics-go-app/internal/api"
	otel "inetum.com/metrics-go-app/internal/otel"
)

func main() {
	// Singleton exit.
	defer func() { _ = otel.TraceProvider.Shutdown(context.Background()) }()
	defer func() { _ = otel.MeterProvider.Shutdown(context.Background()) }()
	defer func() { _ = otel.LoggerProvider.Shutdown(context.Background()) }()

	slog.SetDefault(otel.Logger)

	// Start the API.
	r := api.Router()
	_ = r.Run(":8080")
}
