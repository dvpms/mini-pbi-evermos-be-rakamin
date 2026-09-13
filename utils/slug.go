package utils

import (
	"regexp"
	"strings"
)

var (
	nonAlphaNumericRegex = regexp.MustCompile(`[^a-z0-9]+`)
	trailingDashRegex    = regexp.MustCompile(`^-+|-+$`)
)

// GenerateSlug generates a clean, URL-safe slug from a string
func GenerateSlug(title string) string {
	slug := strings.ToLower(title)
	slug = strings.TrimSpace(slug)
	slug = nonAlphaNumericRegex.ReplaceAllString(slug, "-")
	slug = trailingDashRegex.ReplaceAllString(slug, "")
	return slug
}
