package routers

import (
	"github.com/gin-gonic/gin"
)

// SetupRoutes - setup routes for server
func SetupRoutes() *gin.Engine {
	router := gin.Default()
	router.Use(gin.Logger())
	router.Use(gin.Recovery())

	v1 := router.Group("/api/v1")
	//addCartRoutes(v1)
	addProductRoutes(v1)

	return router
}
