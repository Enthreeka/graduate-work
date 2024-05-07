package utils

import (
	"strings"
)

func KeyGenerator(url string, query string) string {
	var builder strings.Builder

	builder.WriteString(url)
	builder.WriteString("+")
	builder.WriteString(query)

	return builder.String()
}
