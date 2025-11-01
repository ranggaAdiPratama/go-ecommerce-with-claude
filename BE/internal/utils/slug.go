package utils

import (
	"regexp"
	"strings"
)

func GenerateSlug(text string) string {
	slug := strings.ToLower(text)

	reg := regexp.MustCompile(`[^a-z0-9\s-]+`)
	slug = reg.ReplaceAllString(slug, "")

	reg = regexp.MustCompile(`[\s-]+`)
	slug = reg.ReplaceAllString(slug, "-")

	slug = strings.Trim(slug, "-")

	return slug
}
