package config

import (
	"os"

	"github.com/joho/godotenv"
)

type Config struct {
	Port        string `json:"port,omitempty"`
	OllamaURL   string `json:"ollama_url,omitempty"`
	OllamaModel string `json:"ollama_model,omitempty"`
	SupabaseURL string `json:"supabase_url,omitempty"`
	SupabaseKey string `json:"supabase_key,omitempty"`
}

func LoadConfig() *Config {
	_ = godotenv.Load()

	return &Config{
		Port:        getEnv("PORT", "8080"),
		OllamaURL:   getEnv("OLLAMA_URL", "http://localhost:11434"),
		OllamaModel: getEnv("OLLAMA_MODEL", "tinyllama"),
		SupabaseURL: getEnv("SUPABASE_URL", "https://your-supabase-url.supabase.co"),
		SupabaseKey: getEnv("SUPABASE_KEY", "your-supabase-key"),
	}
}

func getEnv(key, defaultValue string) string {
	if value, exists := os.LookupEnv(key); exists {
		return value
	}
	return defaultValue
}
