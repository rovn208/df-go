package products

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func AddNewProduct(c *gin.Context) {
	c.JSON(http.StatusOK, "Create a new product")
}

func GetProducts(c *gin.Context) {
	c.JSON(http.StatusOK, "List of all productRoutes")
}

func UpdateProduct(c *gin.Context) {
	c.JSON(http.StatusOK, "Update product details")
}

func DeleteProduct(c *gin.Context) {
	c.JSON(http.StatusOK, "Delete a product by its ID")
}
