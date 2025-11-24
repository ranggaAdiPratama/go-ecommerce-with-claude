package utils

import (
	"regexp"
	"strings"
	"time"

	"github.com/go-playground/validator/v10"
)

func CapitalizeFirst(s string) string {
	if len(s) == 0 {
		return s
	}
	return strings.ToUpper(s[:1]) + s[1:]
}

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

func Paginator(total int64, limit int32, page int32) (int64, int32) {
	totalPages := (total + int64(limit) - 1) / int64(limit)

	currentPage := int32((page / limit) + 1)

	return totalPages, currentPage
}

func ParseDuration(s string) time.Duration {
	d, err := time.ParseDuration(s)

	if err != nil {
		return 15 * time.Minute
	}

	return d
}
