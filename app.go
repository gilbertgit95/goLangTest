// You can edit this code!
// Click here and start typing.
//  start a godoc in the application
package main

import (
	displayService "golangtest/testservice"

	"github.com/gin-gonic/gin"
)

var Router *gin.Engine

func Plus(x int, y int) int {
	return x + y
}

func Multiply(x int, y int) int {
	return x * y
}

func main() {
	// libService.RunLib()
	displayService.Variables()
	// displayService.Conditions()
	// displayService.Loops()
	// displayService.SliceTypes()
	// displayService.Structs()
	// displayService.Dates()
	// displayService.Interfaces()
	// displayService.Dictionaries()
	// displayService.JSONS()
	// displayService.ConcurChan()

	// test gin rest api
	Router = gin.Default()

	api := Router.Group("/api")
	{
		api.GET("test", func(ctx *gin.Context) {
			ctx.JSON(200, gin.H{"message": "Good"})
		})
	}

	Router.Run(":5000")
}
