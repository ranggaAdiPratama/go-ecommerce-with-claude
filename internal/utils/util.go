package utils

import (
	"regexp"
	"time"

	"github.com/go-playground/validator/v10"
)

func EscapeRegex(s string) string {
	re := regexp.MustCompile(`([\\.^$|(){}\[\]*+?])`)
	return re.ReplaceAllString(s, `\\$1`)
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
