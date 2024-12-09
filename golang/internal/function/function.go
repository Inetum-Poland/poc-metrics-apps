package function

import (
	"context"
	"time"

	"math/rand/v2"

	"go.opentelemetry.io/otel/codes"
	"go.opentelemetry.io/otel/trace"
	otel "inetum.com/metrics-go-app/internal/otel"
	caller "inetum.com/metrics-go-app/internal/utils/caller"
	_recover "inetum.com/metrics-go-app/internal/utils/recover"
)

func randRange(min, max int) int {
	return rand.IntN(max-min) + min
}

func add(ctx context.Context, a int, b int) int {
	callerInfo, _ := caller.GetCallerInfo(1)
	ctx, span := otel.Tracer.Start(ctx, callerInfo.Function, trace.WithAttributes(callerInfo.OtelKV()...))
	defer span.End()
	defer _recover.FuncRecover(ctx, span)

	time.Sleep(time.Millisecond * time.Duration(randRange(20, 800)))

	out := a + b

	span.SetStatus(codes.Ok, "ok")
	return out
}

func substract(ctx context.Context, a int, b int) int {
	callerInfo, _ := caller.GetCallerInfo(1)
	ctx, span := otel.Tracer.Start(ctx, callerInfo.Function, trace.WithAttributes(callerInfo.OtelKV()...))
	defer span.End()
	defer _recover.FuncRecover(ctx, span)

	time.Sleep(time.Millisecond * time.Duration(randRange(20, 800)))

	out := a - b

	span.SetStatus(codes.Ok, "ok")
	return out
}

func multiply(ctx context.Context, a int, b int) int {
	callerInfo, _ := caller.GetCallerInfo(1)
	ctx, span := otel.Tracer.Start(ctx, callerInfo.Function, trace.WithAttributes(callerInfo.OtelKV()...))
	defer span.End()
	defer _recover.FuncRecover(ctx, span)

	time.Sleep(time.Millisecond * time.Duration(randRange(20, 800)))

	out := a * b

	span.SetStatus(codes.Ok, "ok")
	return out
}

func divide(ctx context.Context, a int, b int) int {
	callerInfo, _ := caller.GetCallerInfo(1)
	ctx, span := otel.Tracer.Start(ctx, callerInfo.Function, trace.WithAttributes(callerInfo.OtelKV()...))
	defer span.End()
	defer _recover.FuncRecover(ctx, span)

	time.Sleep(time.Millisecond * time.Duration(randRange(20, 800)))

	out := a / b

	span.SetStatus(codes.Ok, "ok")
	return out
}
