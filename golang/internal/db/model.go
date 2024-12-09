package db

import (
	"inetum.com/metrics-go-app/internal/mongo_orm"
)

type Data struct {
	mongo_orm.Model
	Data int `bson:"data"`
}
