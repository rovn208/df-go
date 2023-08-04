package server

import (
	"github.com/gin-gonic/gin"
)

func New() *gin.Engine {
	router := gin.Default()
	// Product management - Auth BasicAuth/JWT

	// Shopping Cart - Without Auth
	return router
}
