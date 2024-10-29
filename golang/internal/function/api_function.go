package function

import (
	"errors"
	"fmt"
	"log/slog"
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
	span.AddEvent("LongRun started")
	slog.InfoContext(ctx, "LongRun started", slog.String("test", "test"))

	_ = add(ctx, 1, 2)
	_ = substract(ctx, 1, 2)
	_ = multiply(ctx, 1, 2)
	_ = divide(ctx, 1, 2)

	span.SetStatus(codes.Ok, "ok")
	span.AddEvent("LongRun done")
	slog.InfoContext(ctx, "LongRun done", slog.String("test", "test2"))
	c.JSON(http.StatusOK, gin.H{"data": "ok"})
}

func ShortRun(c *gin.Context) {
	_, span := otel.Tracer.Start(c.Request.Context(), "ShortRun")
	defer span.End()
	span.AddEvent("ShortRun started")

	time.Sleep(time.Millisecond * 100)

	span.SetStatus(codes.Ok, "ok")
	span.AddEvent("ShortRun done")
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

	span.AddEvent("DatabaseRun started")

	collection, _, _ := db.Connect(ctx, connectionString)

	span.AddEvent("collection.Find")
	span.SetAttributes(attribute.String("query", fmt.Sprintf("%+v", opts)))

	data, err := db.ReadData(ctx, collection, opts)

	if err != nil {
		fmt.Println(err)
		span.SetStatus(codes.Error, err.Error())
		span.RecordError(err)
		return
	}

	span.SetStatus(codes.Ok, "ok")
	span.AddEvent("DatabaseRun done")
	c.JSON(http.StatusOK, data)
}

func FailedRun(c *gin.Context) {
	_, span := otel.Tracer.Start(c.Request.Context(), "FailedRun")
	defer span.End()
	span.AddEvent("FailedRun started")

	time.Sleep(time.Millisecond * 100)

	fmt.Println("Failed")

	span.SetStatus(codes.Error, "Unexpected error")
	span.RecordError(errors.New("unexpected error"))

	c.JSON(http.StatusInternalServerError, gin.H{"data": "nok"})
}
