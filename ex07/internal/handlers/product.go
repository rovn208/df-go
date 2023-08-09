package handlers

import (
	"github.com/gin-gonic/gin"
	"github.com/rovn208/df-go/ex07/internal/constant"
	"github.com/rovn208/df-go/ex07/internal/model"
	"github.com/rovn208/df-go/ex07/internal/repo"
	"github.com/rovn208/df-go/ex07/internal/util"
)

func AddNewProduct(c *gin.Context) {
	var p model.Product
	if err := c.ShouldBindJSON(&p); err != nil {
		util.BindJSONBadRequest(c, err)
		return
	}
	if err := repo.AddNewProduct(p); err != nil {
		util.BindJSONBadRequest(c, err)
		return
	}

	util.BindSuccessMessageRequest(c, "Product created successfully")
}

func GetProducts(c *gin.Context) {
	products, err := repo.GetProducts()
	if err != nil {
		util.BindJSONBadRequest(c, err)
		return
	}
	util.BindSuccessRequest(c, products)
}

func UpdateProduct(c *gin.Context) {
	var p model.Product
	var productUri model.ProductUri

	if err := c.ShouldBindUri(&productUri); err != nil {
		util.BindJSONBadRequest(c, err)
		return
	}

	if err := c.ShouldBindJSON(&p); err != nil {
		util.BindJSONBadRequest(c, err)
		return
	}

	if p.ID != productUri.ProductId {
		util.BindJSONBadRequest(c, constant.InvalidProductIdError)
		return
	}

	if err := repo.UpdateProduct(p); err != nil {
		util.BindJSONBadRequest(c, err)
		return
	}

	util.BindSuccessRequest(c, p.ID)
}

func DeleteProduct(c *gin.Context) {
	var productUri model.ProductUri

	if err := c.ShouldBindUri(&productUri); err != nil {
		util.BindJSONBadRequest(c, err)
		return
	}

	if err := repo.DeleteProduct(productUri.ProductId); err != nil {
		util.BindJSONBadRequest(c, err)
		return
	}

	util.BindSuccessMessageRequest(c, "product deleted successfully")
}
