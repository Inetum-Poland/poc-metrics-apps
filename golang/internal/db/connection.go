package db

import (
	"context"
	"fmt"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readpref"
	"go.opentelemetry.io/contrib/instrumentation/go.mongodb.org/mongo-driver/mongo/otelmongo"
	"go.opentelemetry.io/otel/codes"
	"go.opentelemetry.io/otel/metric"
	"go.opentelemetry.io/otel/trace"
	"inetum.com/metrics-go-app/internal/otel"
	"inetum.com/metrics-go-app/internal/utils/caller"
)

// "mongodb://root:Password123@mongodb:27017"
func Connect(ctx context.Context, connectionString string) (*mongo.Collection, *mongo.Client, error) {
	// Runtime + Trace + Metric + Log
	callerInfo, _ := caller.GetCallerInfo(1)
	ctx, span := otel.Tracer.Start(ctx, callerInfo.Function, trace.WithAttributes(callerInfo.OtelKV()...))
	otel.DbCounter.Add(ctx, 1, metric.WithAttributes(callerInfo.OtelKV()...))
	span.AddEvent(fmt.Sprintf("%s %s", callerInfo.Function, "started"))
	otel.Logger.InfoContext(ctx, fmt.Sprintf("%s %s", callerInfo.Function, "started"))
	defer span.End()

	mongoOpt := options.Client()
	mongoOpt.ApplyURI(connectionString)
	mongoOpt.Monitor = otelmongo.NewMonitor(otelmongo.WithCommandAttributeDisabled(false))

	client, err := mongo.Connect(ctx, mongoOpt)
	if err != nil {
		fmt.Println(err)
		span.AddEvent("databaseRun error")
		span.SetStatus(codes.Error, err.Error())
		span.RecordError(err)
		return nil, nil, err
	}

	err = client.Ping(ctx, readpref.Primary())
	if err != nil {
		fmt.Println(err)
		span.AddEvent("databaseRun error")
		span.SetStatus(codes.Error, err.Error())
		span.RecordError(err)
		return nil, nil, err
	}

	collection := client.Database("Data").Collection("Data")

	span.SetStatus(codes.Ok, "ok")
	span.AddEvent("Connect done")
	return collection, client, nil
}
