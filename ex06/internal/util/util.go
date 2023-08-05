package util

import (
	"github.com/gin-gonic/gin"
	"github.com/rovn208/df-go/ex06/internal/constants"
	"github.com/rovn208/df-go/ex06/internal/models/carts"
	"github.com/rovn208/df-go/ex06/internal/models/products"
	"net/http"
)

// BindJSONBadRequest binds bad request response
func BindJSONBadRequest(c *gin.Context, err error) {
	c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
}

// BindSuccessMessageRequest binds success response with json message
func BindSuccessMessageRequest(c *gin.Context, msg any) {
	c.JSON(http.StatusOK, gin.H{"message": msg})
}

// BindSuccessRequest binds success response
func BindSuccessRequest(c *gin.Context, msg any) {
	c.JSON(http.StatusOK, msg)
}

// IsProductExists checks if product exists
func IsProductExists(products []products.Product, product products.Product) bool {
	for _, p := range products {
		if p.ID == product.ID {
			return true
		}
	}
	return false
}

// RemoveProductCartAt removes product from cart at index i, throws error if index is invalid
func RemoveProductCartAt(pcs []carts.ProductCart, i int) ([]carts.ProductCart, error) {
	if len(pcs) == 0 || i > len(pcs) {
		return []carts.ProductCart{}, constants.InvalidIndexError
	}
	if i == 0 {
		return pcs[1:], nil
	}
	copyPcs := make([]carts.ProductCart, len(pcs))
	copy(copyPcs, pcs)
	copyPcs[i] = copyPcs[len(copyPcs)-1]
	return copyPcs[:len(copyPcs)-1], nil
}

// RemoveProductAt removes product at index i, throws error if index is invalid
func RemoveProductAt(productList []products.Product, i int) ([]products.Product, error) {
	if len(productList) == 0 || i > len(productList) {
		return []products.Product{}, constants.InvalidIndexError
	}
	if i == 0 {
		return productList[1:], nil
	}
	copyProductList := make([]products.Product, len(productList))
	copy(copyProductList, productList)
	copyProductList[i] = copyProductList[len(copyProductList)-1]
	return copyProductList[:len(copyProductList)-1], nil
}
