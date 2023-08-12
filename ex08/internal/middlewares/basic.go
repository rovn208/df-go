package middlewares

import "github.com/gin-gonic/gin"

// BasicAuth - basic auth middleware
func BasicAuth() gin.HandlerFunc {
	return gin.BasicAuth(gin.Accounts{
		"admin": "admin",
		"dfgo":  "awesome",
	})
}
