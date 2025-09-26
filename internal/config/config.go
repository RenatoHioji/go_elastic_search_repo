package config

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

type Config struct {
	Port     string
	ESUrl    string
	ESUser   string
	ESPass   string
	Postgres string
	RedisUrl string
}

func LoadConfig() Config {
	err := godotenv.Load()
	if err != nil {
		log.Println("No .env file found, using environment variables")
	}
	return Config{
		Port:     getEnv("PORT", "80"),
		ESUrl:    getEnv("ES_URL", "http://localhost:9200"),
		ESUser:   getEnv("ES_USER", ""),
		ESPass:   getEnv("ES_PASS", ""),
		Postgres: getEnv("POSTGRES_URL", "http://localhost:5432"),
		RedisUrl: getEnv("REDIS_URL", "localhost:6379"),
	}
}

func getEnv(key, defaultValue string) string {
	if value := os.Getenv(key); value != "" {
		return value
	}
	return defaultValue
}
