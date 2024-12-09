package main

import (
	"context"
	"fmt"

	"inetum.com/metrics-go-app/internal/api"
	"inetum.com/metrics-go-app/internal/cmd"
	"inetum.com/metrics-go-app/internal/config"
	"inetum.com/metrics-go-app/internal/otel"
)

func main() {
	cmd.RootCommand().Execute()

	otel.Setup()

	defer func() { _ = otel.TraceProvider.Shutdown(context.Background()) }()
	defer func() { _ = otel.MeterProvider.Shutdown(context.Background()) }()
	defer func() { _ = otel.LoggerProvider.Shutdown(context.Background()) }()

	r := api.Router()
	_ = r.Run(fmt.Sprintf(":%d", config.C.WebApp.Port))
}
