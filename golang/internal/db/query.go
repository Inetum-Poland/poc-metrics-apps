package db

import (
	"context"
	"fmt"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.opentelemetry.io/otel/codes"
	"inetum.com/metrics-go-app/internal/otel"
)

func ReadData(ctx context.Context, collection *mongo.Collection, opts bson.M) ([]Data, error) {
	ctx, span := otel.Tracer.Start(ctx, "ReadData")
	defer span.End()
	span.AddEvent("ReadData started")

	findOutput, err := collection.Find(ctx, opts)
	if err != nil {
		fmt.Println(err)
		span.SetStatus(codes.Error, err.Error())
		span.RecordError(err)
		return nil, err
	}

	var data []Data
	if err := findOutput.All(ctx, &data); err != nil {
		fmt.Println(err)
		span.SetStatus(codes.Error, err.Error())
		span.RecordError(err)
		return nil, err
	}

	span.SetStatus(codes.Ok, "ok")
	span.AddEvent("ReadData done")
	return data, nil
}
