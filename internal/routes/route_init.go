package routes

import (
	"log"

	"github.com/RenatoHioji/go_elastic_search_repo/internal/app"
	"github.com/RenatoHioji/go_elastic_search_repo/internal/product"
	"github.com/RenatoHioji/go_elastic_search_repo/internal/search"
	"github.com/RenatoHioji/go_elastic_search_repo/internal/user"
	"github.com/gin-gonic/gin"
)

func InitRoutes(app *app.App) {
	r := gin.Default()

	product.InitProductRoutes(r, app.ProductHandler)
	user.InitUserRoutes(r, app.UserHandler)
	search.InitApiRoutes(r, app.SearchHandler)

	if err := r.Run(app.Config.ServerUrl); err != nil {
		log.Fatal(err)
	}
}
