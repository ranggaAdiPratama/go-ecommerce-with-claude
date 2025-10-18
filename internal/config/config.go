package config

import (
	"ranggaAdiPratama/go-with-claude/internal/utils"
	"time"
)

type Config struct {
	DatabaseURL       string
	Port              string
	Environtment      string
	GinMode           string
	TokenSymmetricKey string
	AccessTokenTTL    time.Duration
	RefreshTokenTTL   time.Duration
}

func Load() *Config {
	return &Config{
		DatabaseURL:       utils.GetEnv("DATABASE_URL", "postgres://rangga:mitsuha@localhost:5432/go?sslmode=disable"),
		Port:              utils.GetEnv("PORT", "8080"),
		Environtment:      utils.GetEnv("ENVIRONMENT", "development"),
		GinMode:           utils.GetEnv("GIN_MODE", "debug"),
		TokenSymmetricKey: utils.GetEnv("TOKEN_SYMMETRIC_KEY", "12345678901234567890123456789012"),
		AccessTokenTTL:    utils.ParseDuration(utils.GetEnv("ACCESS_TOKEN_TTL", "15m")),
		RefreshTokenTTL:   utils.ParseDuration(utils.GetEnv("REFRESH_TOKEN_TTL", "24h")),
	}
}
