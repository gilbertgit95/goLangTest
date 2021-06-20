package controller

import (
	"fmt"
	"io/ioutil"
	"strconv"

	"github.com/gin-gonic/gin"
)

func UserInit(RouterGroup *gin.RouterGroup) {
	fmt.Println("[APP] Initialize User Controller")

	userGroup := RouterGroup.Group("/user")
	{
		// fetch a user
		userGroup.GET("/findone/:userUuid", func(ctx *gin.Context) {
			userUuid := ctx.Param("userUuid")

			ctx.JSON(200, gin.H{
				"message":  "get user",
				"userUuid": userUuid,
			})
		})

		// fetch search all user
		userGroup.GET("/find", func(ctx *gin.Context) {
			fname := ctx.Query("fname")
			mname := ctx.Query("mname")
			lname := ctx.Query("lname")
			role := ctx.Query("role")

			page := ctx.DefaultQuery("page", strconv.Itoa(DefaultPageSkip))
			size := ctx.DefaultQuery("size", strconv.Itoa(DefaultPageSize))

			ctx.JSON(200, gin.H{
				"message": "All users",
				"fname":   fname,
				"lname":   lname,
				"mname":   mname,
				"role":    role,
				"page":    page,
				"size":    size,
			})
		})

		// create users
		userGroup.POST("", func(ctx *gin.Context) {
			body := ctx.Request.Body
			value, err := ioutil.ReadAll(body)

			if err != nil {
				ctx.JSON(500, gin.H{"message": "Error while processing the data."})
			}

			ctx.JSON(200, gin.H{
				"message": "Created Users",
				"data":    string(value),
			})
		})

		// update users
		userGroup.PUT("", func(ctx *gin.Context) {
			ctx.JSON(200, gin.H{"message": "Updated Users"})
		})

		// delete users
		userGroup.DELETE("", func(ctx *gin.Context) {
			ctx.JSON(200, gin.H{"message": "Deleted Users"})
		})
	}
}
