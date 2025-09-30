package product

import (
	"github.com/gin-gonic/gin"
)

func InitProductRoutes(r *gin.Engine, handlers *ProductHandler) {
	group := r.Group("/products")
	{
		group.GET("/", handlers.GetProducts)
		group.POST("/", handlers.CreateProduct)
	}

}
