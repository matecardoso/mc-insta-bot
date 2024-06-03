package utils

import (
	"os"
	"regexp"
)

func GetEnv(key, defaultVal string) string {
	if value, exists := os.LookupEnv(key); exists {
		return value
	}
	return defaultVal
}

func ReplacePlaceholders(input string) string {
	re := regexp.MustCompile(`\{\{(\w+):([^}]+)\}\}`)
	return re.ReplaceAllStringFunc(input, func(match string) string {
		parts := re.FindStringSubmatch(match)
		if len(parts) != 3 {
			return match
		}
		return GetEnv(parts[1], parts[2])
	})
}
