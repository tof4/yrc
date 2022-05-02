package strutil

import (
	"crypto/sha256"
	"fmt"
	"strings"
)

func Sha256(s string) string {
	hash := sha256.Sum256([]byte(s))
	return fmt.Sprintf("%x", hash[:])
}

func Reverse(s string) string {
	runes := []rune(s)
	for i, j := 0, len(runes)-1; i < j; i, j = i+1, j-1 {
		runes[i], runes[j] = runes[j], runes[i]
	}
	return string(runes)
}

func RemoveSpaces(s string) string {
	return strings.ReplaceAll(s, " ", "")
}
