package recover

import (
	"context"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/go-errors/errors"
	"go.opentelemetry.io/otel/attribute"
	"go.opentelemetry.io/otel/codes"
	semconv "go.opentelemetry.io/otel/semconv/v1.26.0"
	"go.opentelemetry.io/otel/trace"
	"inetum.com/metrics-go-app/internal/otel"
)

func anyToAttributes(x []any) []attribute.KeyValue {
	var attrs []attribute.KeyValue

	for _, attr := range x {
		attrs = append(attrs, attr.(attribute.KeyValue))
	}

	return attrs
}

func genericRecover(ctx any, span trace.Span, err any) *errors.Error {
	// In this case '3' is the number of skipped frames.
	goErr := errors.Wrap(err, 3)
	frame := goErr.StackFrames()

	attrs := []any{
		semconv.CodeFilepath(frame[0].File),
		semconv.CodeFunction(frame[0].Func().Name()),
		semconv.CodeLineNumber(frame[0].LineNumber),
		semconv.ExceptionStacktrace(string(goErr.Stack())),
	}

	otel.Logger.ErrorContext(ctx.(context.Context), fmt.Sprintf("%s %s", frame[0].Func().Name(), goErr.Error()), attrs...)

	span.SetStatus(codes.Error, "")
	span.RecordError(errors.New(goErr.Error()),
		trace.WithAttributes(
			anyToAttributes(attrs)...,
		))

	return goErr
}

func GinRecover(c *gin.Context, span trace.Span) {
	if err := recover(); err != nil {
		ctx := c.Request.Context()
		goErr := genericRecover(ctx, span, err)
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": fmt.Sprintf("%v", goErr.Error())})
		c.Error(goErr)
	}
}

func FuncRecover(ctx context.Context, span trace.Span) {
	if err := recover(); err != nil {
		genericRecover(ctx, span, err)
	}
}
