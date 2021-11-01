package app

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"

	"github.com/frankko93/spaceguru-challenge/clients"
	"github.com/frankko93/spaceguru-challenge/src/api/utils"
)

var Router *gin.Engine

func Start() {
	Router.Run(":8080")
}

func init() {
	configureRouter()
	// logger.SetLogLevel(config.LoggerLevel)
	clients.InitDB()
	mapsURLToControllers()
}

func configureRouter() {

	Router = gin.New()
	Router.NoRoute(noRouteHandler)

	Router.Use(utils.HandleRequestID())

	Router.RedirectFixedPath = false
	Router.RedirectTrailingSlash = false
}

func noRouteHandler(c *gin.Context) {
	c.JSON(http.StatusNotFound, fmt.Errorf("Resource %s not found.", c.Request.URL.Path))
}
