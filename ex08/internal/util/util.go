package util

import (
	"github.com/gin-gonic/gin"
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
