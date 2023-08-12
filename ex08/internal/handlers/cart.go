package handlers

import (
	"github.com/gin-gonic/gin"
	"github.com/rovn208/df-go/ex08/internal/model"
	"github.com/rovn208/df-go/ex08/internal/repo"
	"github.com/rovn208/df-go/ex08/internal/util"
)

func AddProduct(c *gin.Context) {
	var pc model.ProductItem

	if err := c.ShouldBindJSON(&pc); err != nil {
		util.BindJSONBadRequest(c, err)
		return
	}
	if err := repo.AddProduct(pc); err != nil {
		util.BindJSONBadRequest(c, err)
		return
	}
	util.BindSuccessMessageRequest(c, "Added a product to cart successfully")
}

func Checkout(c *gin.Context) {
	receipt, err := repo.Checkout()
	if err != nil {
		util.BindJSONBadRequest(c, err)
		return
	}

	util.BindSuccessRequest(c, receipt)
}

func RemoveItem(c *gin.Context) {
	var productIdJson model.ProductIdJson

	if err := c.ShouldBindJSON(&productIdJson); err != nil {
		util.BindJSONBadRequest(c, err)
		return
	}

	if err := repo.RemoveItem(productIdJson.ProductId); err != nil {
		util.BindJSONBadRequest(c, err)
		return
	}

	util.BindSuccessMessageRequest(c, "remove item from carts")
}
