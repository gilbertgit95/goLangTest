package controller

import (
	"fmt"

	"github.com/gin-gonic/gin"
)

func Init(Router *gin.Engine) {
	// api router group
	api := Router.Group("/api")
	{
		// v1 router group under api
		v1 := api.Group("/v1")
		{
			// initialize controllers
			UserInit(v1)
		}

		// v2 router here in the future
		// and so on ...
	}

	fmt.Println("[APP] Initialize Controller")
}
