package config

import (
	"os"
)

type Config struct {
	DatabaseURL string
	Port        string
	Environment string
}

func LoadConfig() (*Config, error) {
	return &Config{
		DatabaseURL: getEnv("DATABASE_URL", ""),
		Port:        getEnv("PORT", ""),
		Environment: getEnv("FLASK_ENV", ""),
	}, nil
}

func getEnv(key, defaultValue string) string {
	if value := os.Getenv(key); value != "" {
		return value
	}
	return defaultValue
}
