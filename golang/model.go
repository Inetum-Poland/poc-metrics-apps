package main

import (
	"context"
	"fmt"

	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readpref"

	"go.opentelemetry.io/contrib/instrumentation/go.mongodb.org/mongo-driver/mongo/otelmongo"
)

type Data struct {
	ID   primitive.ObjectID `bson:"_id"`
	Data int                `bson:"data"`
}

var collection *mongo.Collection

func mongoInit(ctx context.Context) (*mongo.Collection, *mongo.Client) {
	mongoOpt := options.Client()
	mongoOpt.ApplyURI("mongodb://root:Password123@mongodb:27017")
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
