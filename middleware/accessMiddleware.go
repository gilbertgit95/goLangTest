package middleware

import (
	"fmt"

	"github.com/gin-gonic/gin"
)

// this middleware handles the checking of bearer token on the request header
// it then returns 401 if the token is not present or is not valid
func BearerTokenChecker() gin.HandlerFunc {
	return func(ctx *gin.Context) {

		token := ctx.Request.Header.Get("Authorization")

		fmt.Println(token)

		if token == "" {
			ctx.AbortWithStatusJSON(401, gin.H{"error": "API token required"})
		} else {
			ctx.Next()
		}
	}
}

// this function the checking of user roles to insure safety to api that should
// not be accessable by a user
func RoleChecker() gin.HandlerFunc {
	return func(ctx *gin.Context) {

		token := ctx.Request.Header.Get("Authorization")

		fmt.Println(token)

		if token == "" {
			ctx.AbortWithStatusJSON(403, gin.H{"message": "The resource is forbidden for this account type."})
		} else {
			ctx.Next()
		}
	}
}
