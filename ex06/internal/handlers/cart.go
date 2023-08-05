package handlers

import (
	"github.com/gin-gonic/gin"
	"github.com/rovn208/df-go/ex06/internal/db"
	"github.com/rovn208/df-go/ex06/internal/models/carts"
	"github.com/rovn208/df-go/ex06/internal/models/products"
	"github.com/rovn208/df-go/ex06/internal/util"
)

func AddProduct(c *gin.Context) {
	var pc carts.ProductCart

	if err := c.ShouldBindJSON(&pc); err != nil {
		util.BindJSONBadRequest(c, err)
		return
	}
	if err := db.MockDB.AddProduct(pc); err != nil {
		util.BindJSONBadRequest(c, err)
		return
	}
	util.BindSuccessMessageRequest(c, "Added a product to cart successfully")
}

func Checkout(c *gin.Context) {
	receipt, err := db.MockDB.Checkout()
	if err != nil {
		util.BindJSONBadRequest(c, err)
		return
	}

	util.BindSuccessRequest(c, receipt)
}

func RemoveItem(c *gin.Context) {
	var productUri products.ProductUri

	if err := c.ShouldBindJSON(&productUri); err != nil {
		util.BindJSONBadRequest(c, err)
		return
	}

	if err := db.MockDB.RemoveItem(productUri.ProductId); err != nil {
		util.BindJSONBadRequest(c, err)
		return
	}

	util.BindSuccessMessageRequest(c, "remove item from carts")
}
