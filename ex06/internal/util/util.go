package util

import (
	"github.com/gin-gonic/gin"
	"github.com/rovn208/df-go/ex06/internal/constants"
	"github.com/rovn208/df-go/ex06/internal/models/carts"
	"github.com/rovn208/df-go/ex06/internal/models/products"
	"net/http"
)

func BindJSONBadRequest(c *gin.Context, err error) {
	c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
}

func BindSuccessRequest(c *gin.Context, msg any) {
	c.JSON(http.StatusOK, gin.H{"message": msg})

}

func IsProductExists(products []products.Product, product products.Product) bool {
	for _, p := range products {
		if p.ID == product.ID {
			return true
		}
	}
	return false
}

func RemoveProductCartAt(pcs []carts.ProductCart, i int) ([]carts.ProductCart, error) {
	if len(pcs) == 0 || i > len(pcs) {
		return []carts.ProductCart{}, constants.InvalidIndexError
	}
	var copyPcs []carts.ProductCart
	copy(copyPcs, pcs)
	copyPcs[i] = copyPcs[len(copyPcs)-1]
	return copyPcs[:len(copyPcs)-1], nil
}
