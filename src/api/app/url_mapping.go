package app

import (
	"github.com/frankko93/spaceguru-challenge/src/api/controllers"
	"github.com/frankko93/spaceguru-challenge/src/api/middlewares"
)

func mapsURLToControllers() {

	Router.GET("/ping", controllers.Ping)

	users := Router.Group("/users")
	{
		users.POST("/login", controllers.LoginUser)
		users.POST("", controllers.CreateUser)

		users.Use(middlewares.Authenticate())
		users.POST("/travel", controllers.CreateTravel)
		users.GET("/drivers", controllers.SearchUsersDrivers)
	}

}
