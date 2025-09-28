package routes

import (
	"log"

	"github.com/RenatoHioji/go_elastic_search_repo/internal/app"
	"github.com/gin-gonic/gin"
)

func InitRoutes(app *app.App) {
	r := gin.Default()

	InitProductRoutes(r, app.ProductHandler)
	InitUserRoutes(r, app.UserHandler)

	if err := r.Run(app.Config.ServerUrl); err != nil {
		log.Fatal(err)
	}
}
