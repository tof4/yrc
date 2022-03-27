package server

import (
	"crypto/sha256"
	"fmt"
	"log"
)

func authByPassword(username string, password string) bool {
	user, err := getUserByUsername(username)
	if err != nil {
		log.Println(err)
		return false
	}

	hash := sha256.Sum256([]byte(password))
	hashString := fmt.Sprintf("%x", hash[:])
	return user.password == hashString
}
