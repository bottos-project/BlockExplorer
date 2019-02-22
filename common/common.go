package common

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

//ResponseSuccess build successful info for response
func ResponseSuccess(c *gin.Context, message string, result interface{}) {
	c.JSON(http.StatusOK, gin.H{
		"data":    result,
		"message": message,
		"success": true,
	})
}

//ResponseErr build successful info for error
func ResponseErr(c *gin.Context, message string, err error) {
	c.JSON(http.StatusOK, gin.H{
		"data":    err,
		"message": message,
		"success": false,
	})
}
