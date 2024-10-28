package db

import (
	"context"
	"fmt"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readpref"
	"go.opentelemetry.io/contrib/instrumentation/go.mongodb.org/mongo-driver/mongo/otelmongo"
)

// "mongodb://root:Password123@mongodb:27017"
func Connect(ctx context.Context, connectionString string) (*mongo.Collection, *mongo.Client) {
	mongoOpt := options.Client()
	mongoOpt.ApplyURI(connectionString)
	mongoOpt.Monitor = otelmongo.NewMonitor()
	client, err := mongo.Connect(ctx, mongoOpt)
	if err != nil {
		panic(fmt.Sprintf("Mongo DB Connect issue %s", err))
	}

	err = client.Ping(ctx, readpref.Primary())
	if err != nil {
		panic(fmt.Sprintf("Mongo DB ping issue %s", err))
	}

	collection := client.Database("Data").Collection("Data")
	return collection, client
}
