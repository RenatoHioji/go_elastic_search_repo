package config

import (
	"log"

	"github.com/RenatoHioji/go_elastic_search_repo/internal/routes"
	"github.com/gin-gonic/gin"
)

func InitRoutes(cfg Config) {
	r := gin.Default()

	routes.InitApiRoutes(r)

	if err := r.Run(cfg.ServerURL); err != nil {
		log.Fatal(err)
	}

}
