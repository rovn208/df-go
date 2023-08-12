package routers

import (
	"github.com/gin-gonic/gin"
	"github.com/rovn208/df-go/ex08/internal/handlers"
	"github.com/rovn208/df-go/ex08/internal/middlewares"
)

func addProductRoutes(rg *gin.RouterGroup) {
	productRoutes := rg.Group("/products")
	productRoutes.Use(middlewares.BasicAuth())

	productRoutes.POST("/", handlers.AddNewProduct)
	productRoutes.GET("/", handlers.GetProducts)
	productRoutes.PUT("/:product_id", handlers.UpdateProduct)
	productRoutes.DELETE("/:product_id", handlers.DeleteProduct)
}
