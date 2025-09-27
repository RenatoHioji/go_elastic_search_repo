package main

import (
	"github.com/RenatoHioji/go_elastic_search_repo/internal/app"
	"github.com/RenatoHioji/go_elastic_search_repo/internal/config"
)

func main() {
	cfg := config.LoadConfig()
	db := config.InitDB(cfg)

	app.NewApp(db)
	config.InitES(cfg)
	config.InitRoutes(cfg)
}
