
package config

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

type Config struct {
	AppPort      string
	AppEnv       string
	DatabaseURL  string
	MSG91AuthKey string
	MSG91TemplateID string
	JWTSecret    string
}

func LoadConfig() *Config {
	if err := godotenv.Load(); err != nil {
		log.Println("No .env file found; using environment variables")
	}

	cfg := &Config{
		AppPort:         getEnv("APP_PORT", "8080"),
		AppEnv:          getEnv("APP_ENV", "development"),
		DatabaseURL:     getEnv("DATABASE_URL", ""),
		MSG91AuthKey:    getEnv("MSG91_AUTH_KEY", ""),
		MSG91TemplateID: getEnv("MSG91_TEMPLATE_ID", ""),
		JWTSecret:       getEnv("JWT_SECRET", ""),
	}

	return cfg
}

func getEnv(key, fallback string) string {
	value := os.Getenv(key)
	if value == "" {
		return fallback
	}
	return value
}
