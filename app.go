package main

import (
	"golangtest/controller"
	"golangtest/utility/env"

	"github.com/gin-gonic/gin"
)

var Router *gin.Engine

func main() {
	// get environment variables
	appSettings := env.GetENV()

	// instance of the routes
	Router = gin.Default()

	// initiate all the routes and controllers
	controller.Init(Router)

	// run on the port stated on .env
	Router.Run(":" + appSettings.AppPort)
}
