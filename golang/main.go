package main

import (
	"context"
)

func main() {
	ctx := context.Background()
	traceProvider := setupTracer(ctx)
	defer func() { _ = traceProvider.Shutdown(ctx) }()

	r := router()
	_ = r.Run(":8080")
}
