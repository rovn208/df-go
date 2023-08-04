package server

import (
	"github.com/gin-gonic/gin"
	"github.com/rovn208/df-go/ex06/internal/handlers/carts"
)

/*
POST /carts/add: Add items to the carts. It receives a product ID
and quantity as JSON input.
DELETE /carts/remove: Remove items from the carts. It receives
a product ID as JSON input.
POST /carts/checkout: Checkout and clear the carts. It returns a
receipt with the total price.
*/
func addCartRoutes(rg *gin.RouterGroup) {
	cartsRoutes := rg.Group("/carts")

	cartsRoutes.POST("/add", carts.AddProduct)
	cartsRoutes.POST("/checkout", carts.Checkout)
	cartsRoutes.DELETE("/remove", carts.RemoveItem)
}
