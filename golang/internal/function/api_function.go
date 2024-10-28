package function

import (
	"fmt"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson"
	"go.opentelemetry.io/otel/attribute"
	"go.opentelemetry.io/otel/codes"
	"inetum.com/metrics-go-app/internal/db"
	otel "inetum.com/metrics-go-app/internal/otel"
)

func LongRun(c *gin.Context) {
	ctx, span := otel.Tracer.Start(c.Request.Context(), "LongRun")
	defer span.End()
	span.AddEvent("longRun started")

	_ = add(ctx, 1, 2)
	_ = substract(ctx, 1, 2)
	_ = multiply(ctx, 1, 2)
	_ = divide(ctx, 1, 2)

	span.SetStatus(codes.Ok, "ok")
	span.AddEvent("longRun done")
	c.JSON(http.StatusOK, gin.H{"data": "ok"})
}

func ShortRun(c *gin.Context) {
	_, span := otel.Tracer.Start(c.Request.Context(), "LongRun")
	defer span.End()

	time.Sleep(time.Millisecond * 100)
	c.JSON(http.StatusOK, gin.H{"data": "ok"})
}

func DatabaseRun(c *gin.Context) {
	ctx, span := otel.Tracer.Start(c.Request.Context(), "DatabaseRun")
	defer span.End()

	// ---
	connectionString := "mongodb://root:Password123@mongodb:27017"
	opts := bson.M{
		"data": bson.M{
			"$gte": 0,
		},
	}

	span.AddEvent("databaseRun started")

	collection, _ := db.Connect(ctx, connectionString)

	span.AddEvent("collection.Find")
	span.SetAttributes(attribute.String("query", fmt.Sprintf("%+v", opts)))

	data, err := db.Find(collection, opts)

	if err != nil {
		panic(err)
	}

	span.AddEvent("databaseRun done")
	span.SetStatus(codes.Ok, "ok")
	c.JSON(http.StatusOK, data)
}
