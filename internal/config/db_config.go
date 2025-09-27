package config

import (
	"database/sql"
	"log"

	_ "github.com/lib/pq"
)

func InitDB(cfg Config) *sql.DB {
	connStr := cfg.PGUrl
	db, err := sql.Open("postgres", connStr)
	if err != nil {
		log.Fatal(err)
	}
	return db
}
