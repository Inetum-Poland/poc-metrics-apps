package api

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"inetum.com/metrics-go-app/internal/function"
	"inetum.com/metrics-go-app/internal/otel/middleware"
)

// var config = sloggin.Config{
// 	WithRequestBody:    true,
// 	WithResponseBody:   true,
// 	WithRequestHeader:  true,
// 	WithResponseHeader: true,

// 	WithSpanID:  true,
// 	WithTraceID: true,
// }

// func prepare(r *gin.Engine) *gin.Engine {
// 	return r
// }

func Router() *gin.Engine {
	gin.SetMode(gin.ReleaseMode)
	gin.DisableConsoleColor()

	r := gin.New()

	// https://eblog.fly.dev/faststack.html
	// TODO: Panic recovery from OTEL
	// r.Use(utilsMiddlewareGin.Recovery())
	// r.Use(gin.Recovery())

	// go.opentelemetry.io/contrib/instrumentation/github.com/gin-gonic/gin/otelgin
	// TODO: Custom Traces from OTEL??
	// r.Use(otelgin.Middleware("app"))

	// TODO: Metrics from OTEL
	// r.Use(utilsMiddlewareGin.PrometheusMiddleware())

	// Logs from OTEL
	// r.Use(sloggin.NewWithConfig(otel.Logger, config))

	// Middlewares
	r.Use(middleware.Traces())
	r.Use(middleware.Metrics())
	r.Use(middleware.Logs())

	api := r.Group("/api")
	api.GET("/long_run", function.LongRun)
	api.GET("/short_run", function.ShortRun)
	api.GET("/database_run", function.DatabaseRun)
	api.GET("/failed_run", function.FailedRun)

	r.GET("/health", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{"data": "ok"})
	})

	// Default error handler
	r.NoRoute(func(c *gin.Context) {
		c.JSON(http.StatusNotFound, gin.H{"code": "PAGE_NOT_FOUND", "message": "404 page not found"})
	})

	r.NoMethod(func(c *gin.Context) {
		c.JSON(http.StatusMethodNotAllowed, gin.H{"code": "METHOD_NOT_ALLOWED", "message": "405 method not allowed"})
	})

	return r
}
