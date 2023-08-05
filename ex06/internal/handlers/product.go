package handlers

import (
	"github.com/gin-gonic/gin"
	"github.com/rovn208/df-go/ex06/internal/constants"
	"github.com/rovn208/df-go/ex06/internal/db"
	"github.com/rovn208/df-go/ex06/internal/models/products"
	"github.com/rovn208/df-go/ex06/internal/util"
)

func AddNewProduct(c *gin.Context) {
	var p products.Product
	if err := c.ShouldBindJSON(&p); err != nil {
		util.BindJSONBadRequest(c, err)
		return
	}
	if err := db.MockDB.AddNewProduct(p); err != nil {
		util.BindJSONBadRequest(c, err)
		return
	}

	util.BindSuccessMessageRequest(c, "Product created successfully")
}

func GetProducts(c *gin.Context) {
	util.BindSuccessRequest(c, db.MockDB.GetProducts())
}

func UpdateProduct(c *gin.Context) {
	var p products.Product
	var productUri products.ProductUri

	if err := c.ShouldBindUri(&productUri); err != nil {
		util.BindJSONBadRequest(c, err)
		return
	}

	if err := c.ShouldBindJSON(&p); err != nil {
		util.BindJSONBadRequest(c, err)
		return
	}

	if p.ID != productUri.ProductId {
		util.BindJSONBadRequest(c, constants.InvalidProductIdError)
		return
	}

	if err := db.MockDB.UpdateProduct(p); err != nil {
		util.BindJSONBadRequest(c, err)
		return
	}

	util.BindSuccessRequest(c, p)
}

func DeleteProduct(c *gin.Context) {
	var productUri products.ProductUri

	if err := c.ShouldBindUri(&productUri); err != nil {
		util.BindJSONBadRequest(c, err)
		return
	}

	if err := db.MockDB.DeleteProduct(productUri.ProductId); err != nil {
		util.BindJSONBadRequest(c, err)
		return
	}

	util.BindSuccessMessageRequest(c, "product deleted successfully")
}
