package app

import (
	"github.com/frankko93/spaceguru-challenge/src/api/controllers"
	"github.com/frankko93/spaceguru-challenge/src/api/middlewares"
	"github.com/gin-gonic/gin"
)

func mapsURLToControllers() {

	Router.Use(gin.Recovery())
	Router.GET("/ping", controllers.Ping)

	users := Router.Group("/users")
	{
		users.POST("/login", controllers.LoginUser)
		users.POST("", controllers.CreateUser)

		users.Use(middlewares.Authenticate())
		users.GET("/drivers", controllers.SearchUsersDrivers)
	}

	travel := Router.Group("/travel")
	{
		travel.POST("", controllers.CreateTravel)
	}

}
