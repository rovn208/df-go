package routers

import (
	"github.com/gin-gonic/gin"
	"github.com/rovn208/df-go/ex06/internal/handlers"
)

func addCartRoutes(rg *gin.RouterGroup) {
	cartsRoutes := rg.Group("/carts")

	cartsRoutes.POST("/add", handlers.AddProduct)
	cartsRoutes.POST("/checkout", handlers.Checkout)
	cartsRoutes.DELETE("/remove", handlers.RemoveItem)
}
