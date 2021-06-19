package main

import (
	"golangtest/controller"

	"github.com/gin-gonic/gin"
)

var Router *gin.Engine

func main() {
	Router = gin.Default()

	controller.Init(Router)

	Router.Run(":5000")
}
