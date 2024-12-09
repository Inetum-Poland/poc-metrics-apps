package middleware

import (
	"github.com/gin-gonic/gin"
	"go.opentelemetry.io/otel/attribute"
	"go.opentelemetry.io/otel/codes"
	semconv "go.opentelemetry.io/otel/semconv/v1.26.0"
	"go.opentelemetry.io/otel/trace"
	"inetum.com/metrics-go-app/internal/otel"
)

// https://github.com/open-telemetry/opentelemetry-go-contrib/blob/main/instrumentation/github.com/gin-gonic/gin/otelgin/gintrace.go
func Traces() gin.HandlerFunc {
	return func(c *gin.Context) {
		ctx, span := otel.Tracer.Start(c.Request.Context(), c.FullPath(),
			// trace.WithAttributes(semconvutil.HTTPServerRequest(service, c.Request)...),
			trace.WithAttributes(semconv.HTTPRoute(c.FullPath())),
			trace.WithSpanKind(trace.SpanKindServer),
		)
		defer span.End()

		c.Request = c.Request.WithContext(ctx)

		// before
		c.Next()
		// after

		attrs := []attribute.KeyValue{
			semconv.HTTPResponseStatusCode(c.Writer.Status()),
			semconv.HTTPRequestMethodOriginal(c.Request.Method),
			semconv.HTTPRoute(c.Request.Host + c.Request.URL.String()),
			semconv.HTTPRequestSize(int(c.Request.ContentLength)),
			semconv.HTTPResponseSize(int(c.Writer.Size())),
		}

		span.SetAttributes(attrs...)

		if c.IsAborted() || len(c.Errors) > 0 {
			span.SetStatus(codes.Error, c.Errors.String())

			for _, err := range c.Errors {
				span.RecordError(err.Err)
			}
		} else {
			span.SetStatus(codes.Ok, "")
		}
	}
}
