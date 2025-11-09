package stringutil

import (
	"strings"
	"unicode"
)

// ToCamelCase converts snake_case or kebab-case strings to CamelCase.
func ToCamelCase(s string) string {
	if s == "" {
		return s
	}

	// Replace dash with underscore for consistency
	s = strings.ReplaceAll(s, "-", "_")

	parts := strings.Split(s, "_")
	for i, part := range parts {
		if len(part) > 0 {
			parts[i] = strings.ToUpper(string(part[0])) + strings.ToLower(part[1:])
		}
	}
	return strings.Join(parts, "")
}

// ToSnakeCase converts CamelCase or PascalCase strings to snake_case.
func ToSnakeCase(s string) string {
	var result []rune
	for i, r := range s {
		if unicode.IsUpper(r) {
			if i > 0 {
				result = append(result, '_')
			}
			result = append(result, unicode.ToLower(r))
		} else {
			result = append(result, r)
		}
	}
	return string(result)
}

// ContainsIgnoreCase checks if a string contains another string, ignoring case.
func ContainsIgnoreCase(s, substr string) bool {
	return strings.Contains(strings.ToLower(s), strings.ToLower(substr))
}

// Reverse reverses a given string.
func Reverse(s string) string {
	runes := []rune(s)
	for i, j := 0, len(runes)-1; i < j; i, j = i+1, j-1 {
		runes[i], runes[j] = runes[j], runes[i]
	}
	return string(runes)
}
