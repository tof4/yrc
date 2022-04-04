package server

import (
	"fmt"
	"regexp"
	"strings"
)

func getStringsBetweenQuotes(source string) []string {
	r := regexp.MustCompile(`(\w+)||((?:\\"|[^"])*)`)
	match := r.FindAllString(source, -1)
	if len(match) < 2 {
		return nil
	}

	var results []string
	for _, s := range match {
		preparedString := strings.TrimSpace(strings.ReplaceAll(s, `\"`, ""))
		if len(preparedString) > 0 {
			results = append(results, preparedString)
		}
	}

	fmt.Println(results)
	return results
}

func validateMessage(source string) bool {
	if len(source) > 0 && len(source) < 500 {
		return true
	}
	return false
}
