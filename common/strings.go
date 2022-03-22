package common

import (
	"regexp"
	"strings"
)

func GetStringBetweenQuotes(source string) string {
	r, _ := regexp.Compile("'(.*?)'")
	return strings.TrimSpace(r.FindStringSubmatch(source)[1])
}
