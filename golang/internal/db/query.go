package db

import (
	"context"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

func Find(collection *mongo.Collection, opts bson.M) ([]Data, error) {
	findOutput, err := collection.Find(context.Background(), opts)
	if err != nil {
		return nil, err
	}

	var data []Data
	if err := findOutput.All(context.Background(), &data); err != nil {
		return nil, err
	}

	return data, nil
}
