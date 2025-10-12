package config

import (
	"os"

	"github.com/joho/godotenv"
)

type Config struct {
	Port		string
	DatabaseURL	string
	OllamaURL	string
}

func LoadConfig() *Config {
	_ = godotenv.Load()

	return &Config{
		Port:        getEnv("PORT", "8080"),
		DatabaseURL: getEnv("DATABASE_URL", "postgres://user:password@localhost:5432/dbname?sslmode=disable"),
		OllamaURL:   getEnv("OLLAMA_URL", "http://localhost:11434"),
	}
}

func getEnv(key, defaultValue string) string {
	if value, exists := os.LookupEnv(key); exists {
		return value
	}
	return defaultValue
}