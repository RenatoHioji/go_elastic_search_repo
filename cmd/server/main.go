package main

import (
	"log"

	"github.com/RenatoHioji/go_elastic_search_repo/internal/app"
	"github.com/RenatoHioji/go_elastic_search_repo/internal/routes"
	"github.com/RenatoHioji/go_elastic_search_repo/internal/wire"
	"github.com/gin-gonic/gin"
)

func InitRoutes(app *app.App) {
	r := gin.Default()

	routes.InitProductRoutes(r, app.ProductHandler)

	if err := r.Run(app.Config.ServerUrl); err != nil {
		log.Fatal(err)
	}
}

func main() {
	application, err := wire.InitializeApp()
	if err != nil {
		log.Fatal(err)
	}
	InitRoutes(application)
}
