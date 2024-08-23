package config

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

type Config struct {
	PORT   string
	DB_URL string
	// DB_HOST     string
	// DB_PORT     string
	// DB_USER     string
	// DB_PASSWORD string
	// DB_NAME     string
}

var AppConfig Config

func LoadConfig() {
	// Load .env.prod file
	err := godotenv.Load(".env.prod")
	if err != nil {
		log.Fatalf("Error loading .env.prod file: %v", err)
	}

	AppConfig = Config{
		PORT:   getEnv("PORT", "8080"),
		DB_URL: getEnv("DB_URL", ""),
		// DB_HOST:     getEnv("DB_HOST", "localhost"),
		// DB_PORT:     getEnv("DB_PORT", "5432"),
		// DB_USER:     getEnv("DB_USER", "solution"),
		// DB_PASSWORD: getEnv("DB_PASSWORD", "admin"),
		// DB_NAME:     getEnv("DB_NAME", "coody"),
	}
}

// Helper function to get environment variables with a fallback
func getEnv(key, fallback string) string {
	if value, exists := os.LookupEnv(key); exists {
		return value
	}
	return fallback
}
