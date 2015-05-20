package domain

import (
	"github.com/eduncan911/sanitize"
)

// Slug takes a string and converts it into an URL-safe slug
func Slug(s string) string {
	return sanitize.Path(s)
}
