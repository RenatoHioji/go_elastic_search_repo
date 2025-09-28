package routes

import (
	"github.com/RenatoHioji/go_elastic_search_repo/internal/handlers"
	"github.com/gin-gonic/gin"
)

func InitUserRoutes(r *gin.Engine, handlers *handlers.UserHandler) {
	group := r.Group("/users")
	{
		group.GET("/", handlers.GetUsers)
	}
}
