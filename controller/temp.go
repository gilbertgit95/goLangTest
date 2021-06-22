package controller

import (
	"fmt"
	"golangtest/utility/env"

	"github.com/gin-gonic/gin"
)

func TempInit(RouterGroup *gin.RouterGroup) {
	fmt.Println("[APP] Initialize Temp Controller")

	tempGroup := RouterGroup.Group("/temp")
	{
		// return temp route output
		tempGroup.GET("/test", func(ctx *gin.Context) {
			appSettings := env.GetENV()

			ctx.JSON(200, gin.H{
				"appSettings": appSettings,
			})
		})
	}
}
