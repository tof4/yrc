package database

import (
	"os"
	"strings"
	"unicode/utf8"

	"github.com/tof4/yrc/internal/errutil"
	"github.com/tof4/yrc/internal/strutil"
)

func fileAppend(newLine string, filepath string) {
	file, err := os.OpenFile(filepath, os.O_APPEND|os.O_WRONLY|os.O_CREATE, 0600)
	defer file.Close()
	_, err = file.WriteString(newLine)
	errutil.CatchFatal(err)
}

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
				lines[currentLine] = strutil.Reverse(sb.String())
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
