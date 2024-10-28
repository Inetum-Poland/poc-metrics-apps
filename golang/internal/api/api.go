package api

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"go.opentelemetry.io/contrib/instrumentation/github.com/gin-gonic/gin/otelgin"
	apifunc "inetum.com/metrics-go-app/internal/function"
)

// ---

func Router() *gin.Engine {
	gin.SetMode(gin.ReleaseMode)
	r := gin.Default()

	r.Use(gin.Recovery())
	r.Use(otelgin.Middleware("app"))

	api := r.Group("/api")
	{
		api.GET("/long_run", apifunc.LongRun)
		api.GET("/short_run", apifunc.ShortRun)
		api.GET("/database_run", apifunc.DatabaseRun)
	}

	r.GET("/", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{"data": "hello world"})
	})

	return r
}
