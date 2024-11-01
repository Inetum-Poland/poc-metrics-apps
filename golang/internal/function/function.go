package function

import (
	"context"
	"time"

	"math/rand/v2"

	"go.opentelemetry.io/otel/codes"
	otel "inetum.com/metrics-go-app/internal/otel"
)

func randRange(min, max int) int {
	return rand.IntN(max-min) + min
}

func performOperation(ctx context.Context, operation string, a, b int, opFunc func(int, int) int) int {
	_, span := otel.Tracer.Start(ctx, operation)
	defer span.End()
	span.AddEvent(operation + " started")

	time.Sleep(time.Millisecond * time.Duration(randRange(20, 800)))

	out := opFunc(a, b)

	span.SetStatus(codes.Ok, "ok")
	span.AddEvent(operation + " done")
	return out
}

func add(ctx context.Context, a int, b int) int {
	return performOperation(ctx, "Add", a, b, func(a, b int) int { return a + b })
}

func substract(ctx context.Context, a int, b int) int {
	return performOperation(ctx, "Substract", a, b, func(a, b int) int { return a - b })
}

func multiply(ctx context.Context, a int, b int) int {
	return performOperation(ctx, "Multiply", a, b, func(a, b int) int { return a * b })
}

func divide(ctx context.Context, a int, b int) int {
	return performOperation(ctx, "Multiply", a, b, func(a, b int) int { return a / b })
}
