package middleware

import (
	"log/slog"

	"github.com/gin-gonic/gin"
	semconv "go.opentelemetry.io/otel/semconv/v1.26.0"
	"inetum.com/metrics-go-app/internal/otel"
)

// https://github.com/samber/slog-gin/blob/main/middleware.go
func Logs() gin.HandlerFunc {
	return func(c *gin.Context) {

		// before
		c.Next()
		// after

		attrs := []any{
			slog.Int(string(semconv.HTTPResponseStatusCodeKey), c.Writer.Status()),
			slog.String(string(semconv.HTTPRequestMethodKey), c.Request.Method),
			slog.String(string(semconv.HTTPRouteKey), c.Request.Host+c.Request.URL.String()),
			slog.Int(string(semconv.HTTPRequestSizeKey), int(c.Request.ContentLength)),
			slog.Int(string(semconv.HTTPResponseSizeKey), int(c.Writer.Size())),

			// semconv.HTTPResponseStatusCode(c.Writer.Status()),
			// semconv.HTTPRequestMethodOriginal(c.Request.Method),
			// semconv.HTTPRoute(c.Request.Host + c.Request.URL.String()),
			// attribute.Int("http.request.size", int(c.Request.ContentLength)),
			// attribute.Int("http.response.size", int(c.Writer.Size())),
			// semconv.HTTPRequestSize(int(c.Request.ContentLength)),
			// semconv.HTTPResponseSize(int(c.Writer.Size())),
			// semconv.HTTPRequestBodySize(),
			// semconv.HTTPResponseBodySize(),
		}

		if c.IsAborted() || len(c.Errors) > 0 {
			otel.Logger.ErrorContext(c.Request.Context(), c.Request.RequestURI, attrs...)
		} else {
			otel.Logger.InfoContext(c.Request.Context(), c.Request.RequestURI, attrs...)
		}
	}
}
