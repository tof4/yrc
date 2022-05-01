package common

import (
	"crypto/sha256"
	"fmt"
)

func Sha256String(input string) string {
	hash := sha256.Sum256([]byte(input))
	return fmt.Sprintf("%x", hash[:])
}
