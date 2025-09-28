package user

import (
	"github.com/gin-gonic/gin"
)

func InitUserRoutes(r *gin.Engine, handlers *UserHandler) {
	group := r.Group("/users")
	{
		group.GET("/", handlers.GetUsers)
	}
}
