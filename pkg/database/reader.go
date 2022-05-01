package database

import (
	"os"
	"strings"
	"unicode/utf8"
)

func BackwardFileRead(path string, count int) []string {
	file, _ := os.Open(path)
	defer file.Close()

	buf := make([]byte, 1)
	lines := make([]string, count)
	var sb strings.Builder
	start, _ := file.Seek(0, 2)

	currentLine := count - 1

	for i := start; i >= 0; i-- {
		if currentLine == -1 {
			break
		}

		file.ReadAt(buf, i)

		c, _ := utf8.DecodeRune(buf)
		if c == utf8.RuneError && len(buf) <= 5 {
			buf = make([]byte, len(buf)+1)
		} else {
			if c == '\n' {
				lines[currentLine] = reverse(sb.String())
				sb.Reset()
				buf = make([]byte, 1)
				currentLine--
			} else {
				sb.WriteRune(c)
			}
		}
	}

	return lines
}

func reverse(s string) string {
	runes := []rune(s)
	for i, j := 0, len(runes)-1; i < j; i, j = i+1, j-1 {
		runes[i], runes[j] = runes[j], runes[i]
	}
	return string(runes)
}
