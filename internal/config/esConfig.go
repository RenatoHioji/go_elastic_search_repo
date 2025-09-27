package config

import (
	"log"

	"github.com/elastic/go-elasticsearch/v9"
)

func InitES(cfg Config) {
	_, err := elasticsearch.NewClient(elasticsearch.Config{
		Addresses: []string{cfg.ESUrl},
	})

	if err != nil {
		log.Fatalf("Error creating client: %s", err)
	}
}
