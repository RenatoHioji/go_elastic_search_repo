package main

import (
	"github.com/RenatoHioji/go_elastic_search_repo/internal/config"
)

func main() {
	cfg := config.LoadConfig()

	config.InitES(cfg)
	config.InitRoutes(cfg)
}
