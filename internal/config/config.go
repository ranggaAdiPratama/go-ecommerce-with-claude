package config

import "os"

type Config struct {
	DatabaseURL  string
	Port         string
	Environtment string
	GinMode      string
}

func Load() *Config {
	return &Config{
		DatabaseURL:  getEnv("DATABASE_URL", "postgres://rangga:mitsuha@localhost:5432/go?sslmode=disable"),
		Port:         getEnv("PORT", "8080"),
		Environtment: getEnv("ENVIRONMENT", "development"),
		GinMode:      getEnv("GIN_MODE", "debug"),
	}
}

func getEnv(key, fallback string) string {
	if value := os.Getenv(key); value != "" {
		return value
	}

	return fallback
}
