package server

import (
	"github.com/gin-gonic/gin"
	"github.com/rovn208/df-go/ex06/internal/handlers/products"
	"github.com/rovn208/df-go/ex06/internal/middlewares"
)

/*
POST /products: Create a new product. It receives product
details as JSON input.
GET /products: Retrieve a list of all products.
PUT /products/{product_id}: Update a product's details. It
receives updated product details as JSON input.
DELETE /products/{product_id}: Delete a product by its ID.
*/
func addProductRoutes(rg *gin.RouterGroup) {
	productRoutes := rg.Group("/products")
	productRoutes.Use(middlewares.BasicAuth())

	productRoutes.POST("/", products.AddNewProduct)
	productRoutes.GET("/", products.GetProducts)
	productRoutes.PUT("/:productID", products.UpdateProduct)
	productRoutes.DELETE("/:productID", products.DeleteProduct)
}
