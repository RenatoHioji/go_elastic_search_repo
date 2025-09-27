package config

import (
	"log"

	"github.com/RenatoHioji/go_elastic_search_repo/internal/routes"
	"github.com/gin-gonic/gin"
)

func InitRoutes(cfg Config) {
	r := gin.Default()

	routes.InitApiRoutes(r)
	routes.InitUserRoutes(r)
	routes.InitProductRoutes(r)
	if err := r.Run(cfg.ServerUrl); err != nil {
		log.Fatal(err)
	}

}
