package config

import (
	"os"
	"ranggaAdiPratama/go-with-claude/internal/utils"
	"time"
)

type Config struct {
	DatabaseURL         string
	Port                string
	Environtment        string
	GinMode             string
	TokenSymmetricKey   string
	AccessTokenTTL      time.Duration
	RefreshTokenTTL     time.Duration
	CloudinaryName      string
	CloudinaryAPIKey    string
	CloudinaryAPISecret string
	CloudinaryFolder    string
}

func Load() *Config {
	return &Config{
		DatabaseURL:         GetEnv("DATABASE_URL", "postgres://rangga:mitsuha@localhost:5432/go?sslmode=disable"),
		Port:                GetEnv("PORT", "8080"),
		Environtment:        GetEnv("ENVIRONMENT", "development"),
		GinMode:             GetEnv("GIN_MODE", "debug"),
		TokenSymmetricKey:   GetEnv("TOKEN_SYMMETRIC_KEY", "12345678901234567890123456789012"),
		AccessTokenTTL:      utils.ParseDuration(GetEnv("ACCESS_TOKEN_TTL", "15m")),
		RefreshTokenTTL:     utils.ParseDuration(GetEnv("REFRESH_TOKEN_TTL", "24h")),
		CloudinaryName:      GetEnv("CLOUDINARY_NAME", "thisone"),
		CloudinaryAPIKey:    GetEnv("CLOUDINARY_API_KEY", "234567890"),
		CloudinaryAPISecret: GetEnv("CLOUDINARY_API_SECRET", "mautauajaataumautaubanget?"),
		CloudinaryFolder:    GetEnv("CLOUDINARY_FOLDER", "shops"),
	}
}

func GetEnv(key, fallback string) string {
	if value := os.Getenv(key); value != "" {
		return value
	}

	return fallback
}
