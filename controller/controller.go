package controller

import (
	"fmt"
	middleware "golangtest/middleware"

	"github.com/gin-gonic/contrib/static"
	"github.com/gin-gonic/gin"
)

const (
	DefaultPageSize = 100
	DefaultPageSkip = 0
)

func Init(Router *gin.Engine) {
	fmt.Println("[APP] Initialize Controller")

	// serve static files
	Router.Use(static.Serve("/", static.LocalFile("./webapp/build", true)))

	// api router group
	api := Router.Group("/api")
	{
		// check authorizition token
		api.Use(middleware.BearerTokenChecker())
		// check user roles acces to the route
		api.Use(middleware.RoleChecker())

		// v1 router group under api
		v1 := api.Group("/v1")
		{
			// initialize controllers
			UserInit(v1)
		}

		// v2 router here in the future
		// and so on ...
	}

	// other api group
}
