package main

import (
	"fmt"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson"
	"go.opentelemetry.io/contrib/instrumentation/github.com/gin-gonic/gin/otelgin"
	"go.opentelemetry.io/otel/attribute"
	"go.opentelemetry.io/otel/codes"
)

func longRun(c *gin.Context) {
	ctx, span := tracer.Start(c.Request.Context(), "longRun")
	defer span.End()
	span.AddEvent("longRun started")

	firstFunction(ctx, 1, 2)
	secondFunction(ctx, 1, 2)
	thirdFunction(ctx, 1, 2)
	fourthFunction(ctx, 1, 2)

	span.SetStatus(codes.Ok, "ok")
	span.AddEvent("longRun done")
	c.JSON(http.StatusOK, gin.H{"data": "ok"})
}

func shortRun(c *gin.Context) {
	time.Sleep(time.Millisecond * 100)
	c.JSON(http.StatusOK, gin.H{"data": "ok"})
}

func databaseRun(c *gin.Context) {
	ctx, span := tracer.Start(c.Request.Context(), "databaseRun")
	defer span.End()
	span.AddEvent("databaseRun started")

	collection, _ := mongoInit(ctx)

	opts := bson.M{
		"data": bson.M{
			"$gte": 0,
		},
	}

	findOutput, err := collection.Find(ctx, opts)

	if err != nil {
		panic(err)
	}
	span.AddEvent("collection.Find")
	span.SetAttributes(attribute.String("query", fmt.Sprintf("%+v", opts)))

	var data []Data
	if err := findOutput.All(ctx, &data); err != nil {
		panic(err)
	}

	span.AddEvent("databaseRun done")
	span.SetStatus(codes.Ok, "ok")
	c.JSON(http.StatusOK, data)
}

// ---

func router() *gin.Engine {
	gin.SetMode(gin.ReleaseMode)
	r := gin.Default()
	r.Use(prometheusMiddleware())
	r.Use(otelgin.Middleware("app"))

	api := r.Group("/api")
	{
		api.GET("/long_run", longRun)
		api.GET("/short_run", shortRun)
		api.GET("/database_run", databaseRun)
	}

	r.GET("/", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{"data": "hello world"})
	})

	r.GET("/favicon.ico", func(c *gin.Context) { c.AbortWithStatus(http.StatusOK) })
	r.GET("/metrics", prometheusHandler())
	r.Use(gin.Recovery())

	return r
}
