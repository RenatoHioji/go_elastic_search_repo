package main

import (
	"log"

	"github.com/RenatoHioji/go_elastic_search_repo/internal/routes"
	"github.com/RenatoHioji/go_elastic_search_repo/internal/wire"
)

func main() {
	app, err := wire.InitializeApp()
	if err != nil {
		log.Fatal(err)
	}

	routes.InitRoutes(app)
}
