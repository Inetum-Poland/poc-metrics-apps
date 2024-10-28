package main

import (
	"context"

	api "inetum.com/metrics-go-app/internal/api"
	otel "inetum.com/metrics-go-app/internal/otel"
)

func main() {
	// Singleton exit.
	defer func() { _ = otel.TraceProvider.Shutdown(context.Background()) }()

	// Start the API.
	r := api.Router()
	_ = r.Run(":8080")
}
