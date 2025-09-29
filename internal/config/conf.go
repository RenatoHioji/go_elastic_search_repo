package config

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

type Config struct {
	ServerUrl      string
	TrustedProxies string
	ESUrl          string
	ESUser         string
	ESPass         string
	PGUrl          string
	RedisUrl       string
	KafkaUrl       string
}

func LoadConfig() (Config, error) {
	err := godotenv.Load()
	if err != nil {
		log.Println("No .env file found, using environment variables")
	}
	return Config{
		ServerUrl:      getEnv("ServerUrl", "127.0.0.1:8080"),
		TrustedProxies: getEnv("TrustedProxies", "127.0.0.1"),
		ESUrl:          getEnv("ES_URL", "http://localhost:9200"),
		ESUser:         getEnv("ES_USER", ""),
		ESPass:         getEnv("ES_PASS", ""),
		PGUrl:          getEnv("POSTGRES_URL", "postgres://postgres:root@localhost/elasticsearch?sslmode=disable"),
		RedisUrl:       getEnv("REDIS_URL", "localhost:6379"),
		KafkaUrl:       getEnv("KAFKA_URL", "localhost:9092"),
	}, nil
}

func getEnv(key, defaultValue string) string {
	if value := os.Getenv(key); value != "" {
		return value
	}
	return defaultValue
}
