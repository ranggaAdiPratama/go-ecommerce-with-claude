package utils

import (
	"os"
	"time"

	"github.com/go-playground/validator/v10"
)

func GetEnv(key, fallback string) string {
	if value := os.Getenv(key); value != "" {
		return value
	}

	return fallback
}

func HumanizeError(fe validator.FieldError) string {
	switch fe.Tag() {
	case "required":
		return fe.Field() + " is required"
	case "email":
		return fe.Field() + " is not a valid email"
	case "min":
		return fe.Field() + " must be at least " + fe.Param() + " characters long"
	default:
		return fe.Field() + " is invalid"
	}
}

func ParseDuration(s string) time.Duration {
	d, err := time.ParseDuration(s)

	if err != nil {
		return 15 * time.Minute
	}

	return d
}
