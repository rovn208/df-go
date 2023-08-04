package carts

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func AddProduct(c *gin.Context) {
	c.JSON(http.StatusOK, "add item to carts")
}

func Checkout(c *gin.Context) {
	c.JSON(http.StatusOK, "carts checkout")
}

func RemoveItem(c *gin.Context) {
	c.JSON(http.StatusOK, "remove item from carts")
}
