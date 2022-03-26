package server

import (
	"regexp"
	"strings"
)

func getStringBetweenQuotes(source string) string {
	r, _ := regexp.Compile(`"((?:\\"|[^"])*)"`)
	content := r.FindStringSubmatch(source)[1]
	content = strings.ReplaceAll(content, `\"`, "")
	return strings.TrimSpace(content)
}
