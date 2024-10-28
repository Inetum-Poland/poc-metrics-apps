package db

import (
	"context"
	"fmt"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readpref"
	"go.opentelemetry.io/contrib/instrumentation/go.mongodb.org/mongo-driver/mongo/otelmongo"
	"go.opentelemetry.io/otel/codes"
	"inetum.com/metrics-go-app/internal/otel"
)

// "mongodb://root:Password123@mongodb:27017"
func Connect(ctx context.Context, connectionString string) (*mongo.Collection, *mongo.Client, error) {
	ctx, span := otel.Tracer.Start(ctx, "Connect")
	defer span.End()
	span.AddEvent("Connect started")

	mongoOpt := options.Client()
	mongoOpt.ApplyURI(connectionString)
	mongoOpt.Monitor = otelmongo.NewMonitor()

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
