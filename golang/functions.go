package main

import (
	"context"
	"time"

	"golang.org/x/exp/rand"
)

func firstFunction(ctx context.Context, a int, b int) int {
	_, span := tracer.Start(ctx, "firstFunction")
	defer span.End()

	time.Sleep(time.Millisecond * time.Duration(rand.Intn(1000-200)+200))
	return a + b
}

func secondFunction(ctx context.Context, a int, b int) int {
	_, span := tracer.Start(ctx, "secondFunction")
	defer span.End()

	time.Sleep(time.Millisecond * time.Duration(rand.Intn(1000-200)+200))
	return a - b
}

func thirdFunction(ctx context.Context, a int, b int) int {
	_, span := tracer.Start(ctx, "thirdFunction")
	defer span.End()

	time.Sleep(time.Millisecond * time.Duration(rand.Intn(1000-200)+200))
	return a * b
}

func fourthFunction(ctx context.Context, a int, b int) int {
	ctx, span := tracer.Start(ctx, "fourthFunction")
	defer span.End()

	firstFunction(ctx, a, b)
	secondFunction(ctx, a, b)

	time.Sleep(time.Millisecond * time.Duration(rand.Intn(1000-200)+200))
	return a / b
}
