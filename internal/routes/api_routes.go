package routes

import (
	"github.com/RenatoHioji/go_elastic_search_repo/internal/handlers"
	"github.com/gin-gonic/gin"
)

func InitApiRoutes(r *gin.Engine) {
	group := r.Group("/api")
	{
		group.GET("/", handlers.GetApi)
	}

}
