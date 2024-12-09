package function

import (
	"fmt"
	"log/slog"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson"
	"go.opentelemetry.io/otel/attribute"
	"go.opentelemetry.io/otel/codes"
	"go.opentelemetry.io/otel/metric"
	"go.opentelemetry.io/otel/trace"
	"inetum.com/metrics-go-app/internal/db"
	otel "inetum.com/metrics-go-app/internal/otel"

	// otelLogsGin "inetum.com/metrics-go-app/internal/otel/logs/gin"
	caller "inetum.com/metrics-go-app/internal/utils/caller"
	_recover "inetum.com/metrics-go-app/internal/utils/recover"
)

func LongRun(c *gin.Context) {
	// Runtime + Trace + Metric + Log
	callerInfo, _ := caller.GetCallerInfo(1)
	ctx, span := otel.Tracer.Start(c.Request.Context(), callerInfo.Function, trace.WithAttributes(callerInfo.OtelKV()...))
	defer span.End()
	defer _recover.GinRecover(c, span)

	otel.ApiCounter.Add(ctx, 1, metric.WithAttributes(callerInfo.OtelKV()...))
	span.AddEvent(fmt.Sprintf("%s %s", callerInfo.Function, "started"))
	otel.Logger.InfoContext(ctx, fmt.Sprintf("%s %s", callerInfo.Function, "started"))

	_ = add(ctx, 1, 2)
	_ = substract(ctx, 1, 2)
	_ = multiply(ctx, 1, 2)
	_ = divide(ctx, 1, 2)

	// Trace + Log
	span.SetStatus(codes.Ok, "ok")
	span.AddEvent(fmt.Sprintf("%s %s", callerInfo.Function, "done"))
	otel.Logger.InfoContext(ctx, fmt.Sprintf("%s %s", callerInfo.Function, "done"))

	c.JSON(http.StatusOK, gin.H{"data": "ok"})
}

func ShortRun(c *gin.Context) {
	// Runtime + Trace + Metric + Log
	callerInfo, _ := caller.GetCallerInfo(1)
	ctx, span := otel.Tracer.Start(c.Request.Context(), callerInfo.Function, trace.WithAttributes(callerInfo.OtelKV()...))
	defer span.End()
	defer _recover.GinRecover(c, span)

	otel.ApiCounter.Add(ctx, 1, metric.WithAttributes(callerInfo.OtelKV()...))
	span.AddEvent(fmt.Sprintf("%s %s", callerInfo.Function, "started"))
	otel.Logger.InfoContext(ctx, fmt.Sprintf("%s %s", callerInfo.Function, "started"))

	time.Sleep(time.Millisecond * 100)

	// Trace + Log
	span.SetStatus(codes.Ok, "ok")
	span.AddEvent(fmt.Sprintf("%s %s", callerInfo.Function, "done"))
	otel.Logger.InfoContext(ctx, fmt.Sprintf("%s %s", callerInfo.Function, "done"))

	c.JSON(http.StatusOK, gin.H{"data": "ok"})
}

func DatabaseRun(c *gin.Context) {
	// Runtime + Trace + Metric + Log
	callerInfo, _ := caller.GetCallerInfo(1)
	ctx, span := otel.Tracer.Start(c.Request.Context(), callerInfo.Function, trace.WithAttributes(callerInfo.OtelKV()...))
	defer span.End()
	defer _recover.GinRecover(c, span)

	otel.ApiCounter.Add(ctx, 1, metric.WithAttributes(callerInfo.OtelKV()...))
	span.AddEvent(fmt.Sprintf("%s %s", callerInfo.Function, "started"))
	otel.Logger.InfoContext(ctx, fmt.Sprintf("%s %s", callerInfo.Function, "started"))

	connectionString := "mongodb://root:Password123@mongodb:27017"
	opts := bson.M{
		"data": bson.M{
			"$gte": 0,
		},
	}

	collection, _, _ := db.Connect(ctx, connectionString)

	// Trace + Log
	span.AddEvent("collection.Find")
	span.SetAttributes(attribute.String("query", fmt.Sprintf("%+v", opts)))
	otel.Logger.InfoContext(ctx, "collection.Find", slog.String("query", fmt.Sprintf("%+v", opts)))

	data, err := db.ReadData(ctx, collection, opts)

	if err != nil {
		// Trace + Log
		span.SetStatus(codes.Error, err.Error())
		span.RecordError(err)
		otel.Logger.ErrorContext(ctx, fmt.Sprintf("%s %s", callerInfo.Function, "done"))

		return
	}

	// Trace + Log
	span.SetStatus(codes.Ok, "ok")
	span.AddEvent(fmt.Sprintf("%s %s", callerInfo.Function, "done"))
	otel.Logger.InfoContext(ctx, fmt.Sprintf("%s %s", callerInfo.Function, "done"))

	c.JSON(http.StatusOK, data)
}

func FailedRun(c *gin.Context) {
	// Runtime + Trace + Metric + Log
	callerInfo, _ := caller.GetCallerInfo(1)
	ctx, span := otel.Tracer.Start(c.Request.Context(), callerInfo.Function, trace.WithAttributes(callerInfo.OtelKV()...))
	defer span.End()
	defer _recover.GinRecover(c, span)

	otel.ApiCounter.Add(ctx, 1, metric.WithAttributes(callerInfo.OtelKV()...))
	span.AddEvent(fmt.Sprintf("%s %s", callerInfo.Function, "started"))
	otel.Logger.InfoContext(ctx, fmt.Sprintf("%s %s", callerInfo.Function, "started"))

	// defer otelLogsGin.Recover(c, span)

	time.Sleep(time.Millisecond * 100)

	//! FAIL
	panic("Purpose error")

	// Trace + Log
	// span.SetStatus(codes.Error, "Unexpected error")
	// span.RecordError(errors.New("unexpected error"))
	// otel.Logger.ErrorContext(ctx, fmt.Sprintf("%s %s", callerInfo.Function, "done"))

	// c.JSON(http.StatusInternalServerError, gin.H{"data": "nok"})
}
