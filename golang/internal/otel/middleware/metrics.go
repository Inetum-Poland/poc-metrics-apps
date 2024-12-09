package middleware

import (
	"time"

	"github.com/gin-gonic/gin"
	"go.opentelemetry.io/otel/attribute"
	"go.opentelemetry.io/otel/metric"
	semconv "go.opentelemetry.io/otel/semconv/v1.26.0"
	"inetum.com/metrics-go-app/internal/otel"
)

// https://medium.com/@kylelzk/prometheus-practical-lab-how-to-create-metrics-with-go-client-api-b4119c4f8755
func Metrics() gin.HandlerFunc {
	return func(c *gin.Context) {
		__t := time.Now()

		// before
		c.Next()
		// after

		__latency := time.Since(__t)

		attrs := []attribute.KeyValue{
			semconv.HTTPResponseStatusCode(c.Writer.Status()),
			semconv.HTTPRequestMethodOriginal(c.Request.Method),
			semconv.HTTPRoute(c.Request.Host + c.Request.URL.String()),
			// semconv.HTTPRequestSize(int(c.Request.ContentLength)),
			// semconv.HTTPResponseSize(int(c.Writer.Size())),
		}

		otel.ApiReqCount.Add(c.Request.Context(), 1,
			metric.WithAttributes(attrs...))

		otel.ApiReqTime.Record(c.Request.Context(), __latency.Seconds(),
			metric.WithAttributes(attrs...))
	}
}
