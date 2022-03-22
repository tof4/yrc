package server

import (
	"regexp"
	"strings"
)

func getStringBetweenQuotes(source string) string {
	r, _ := regexp.Compile("'(.*?)'")
	return strings.TrimSpace(r.FindStringSubmatch(source)[1])
}
