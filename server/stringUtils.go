package server

import (
	"regexp"
	"strings"
)

func getStringBetweenQuotes(source string) string {
	r, _ := regexp.Compile(`"((?:\\"|[^"])*)"`)
	match := r.FindStringSubmatch(source)
	if len(match) < 2 {
		return ""
	}

	result := strings.ReplaceAll(match[1], `\"`, "")
	return strings.TrimSpace(result)
}

func validateMessage(source string) bool {
	if len(source) > 0 && len(source) < 500 {
		return true
	}
	return false
}
