package util

import (
	"github.com/gin-gonic/gin"
	"github.com/rovn208/df-go/ex07/internal/constant"
	"github.com/rovn208/df-go/ex07/internal/model"
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
func IsProductExists(products []model.Product, product model.Product) bool {
	for _, p := range products {
		if p.ID == product.ID {
			return true
		}
	}
	return false
}

//
//// RemoveProductCartAt removes product from cart at index i, throws error if index is invalid
//func RemoveProductCartAt(pcs []model.ProductItem, i int) ([]model.ProductItem, error) {
//	if len(pcs) == 0 || i > len(pcs) {
//		return []model.ProductItem{}, constant.InvalidIndexError
//	}
//	if i == 0 {
//		return pcs[1:], nil
//	}
//	copyPcs := make([]model.ProductItem, len(pcs))
//	copy(copyPcs, pcs)
//	copyPcs[i] = copyPcs[len(copyPcs)-1]
//	return copyPcs[:len(copyPcs)-1], nil
//}

// RemoveProductAt removes product at index i, throws error if index is invalid
func RemoveProductAt(productList []model.Product, i int) ([]model.Product, error) {
	if len(productList) == 0 || i > len(productList) {
		return []model.Product{}, constant.InvalidIndexError
	}
	if i == 0 {
		return productList[1:], nil
	}
	copyProductList := make([]model.Product, len(productList))
	copy(copyProductList, productList)
	copyProductList[i] = copyProductList[len(copyProductList)-1]
	return copyProductList[:len(copyProductList)-1], nil
}
