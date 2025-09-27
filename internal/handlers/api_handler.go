package handlers

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func GetApi(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"message": "Hello FROM get API",
	})
}
