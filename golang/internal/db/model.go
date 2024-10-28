package db

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Data struct {
	ID   primitive.ObjectID `bson:"_id"`
	Data int                `bson:"data"`
}
