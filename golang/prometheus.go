package main

import (
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promhttp"
)

var reqCount = prometheus.NewCounterVec(
	prometheus.CounterOpts{
		Name: "req_count",
		Help: "Request count",
	},
	[]string{"method", "path", "status"},
)

var reqTime = prometheus.NewHistogramVec(
	prometheus.HistogramOpts{
		Name:    "req_time",
		Help:    "Request time",
		Buckets: []float64{0.001, 0.005, 0.010, 0.050, 0.100, 0.500, 1.000, 5.000, 10.000},
	},
	[]string{"method", "path", "status"},
)

func init() {
	prometheus.MustRegister(reqCount, reqTime)
}

func prometheusHandler() gin.HandlerFunc {
	h := promhttp.Handler()

	return func(c *gin.Context) {
		h.ServeHTTP(c.Writer, c.Request)
	}
}

// https://medium.com/@kylelzk/prometheus-practical-lab-how-to-create-metrics-with-go-client-api-b4119c4f8755
func prometheusMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		t := time.Now()

		// before

		c.Next()

		// after

		__path := c.Request.URL.Path
		__status := c.Writer.Status()
		__latency := time.Since(t)
		__method := c.Request.Method
		reqCount.WithLabelValues(__method, __path, strconv.Itoa(__status)).Inc()
		reqTime.WithLabelValues(__method, __path, strconv.Itoa(__status)).Observe(__latency.Seconds())
	}
}
