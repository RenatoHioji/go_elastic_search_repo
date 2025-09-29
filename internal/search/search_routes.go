package search

import (
	"github.com/gin-gonic/gin"
)

func InitApiRoutes(r *gin.Engine, handler *SearchHandler) {
	group := r.Group("/search")
	{
		group.GET("/products", handler.SearchProduct)
	}

}
