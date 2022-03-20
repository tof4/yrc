package common

import "regexp"

func GetStringBetween(source string, separator string) string {
	rx := regexp.MustCompile(`(?s)` + regexp.QuoteMeta(separator) + `(.*?)` + regexp.QuoteMeta(separator))
	return rx.FindString(source)
}
