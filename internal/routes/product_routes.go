package routes

import (
	"github.com/RenatoHioji/go_elastic_search_repo/internal/handlers"
	"github.com/gin-gonic/gin"
)

func InitProductRoutes(r *gin.Engine) {
	group := r.Group("/products")
	{
		group.GET("/", handlers.GetProducts)
	}

}
