package server

import (
	"crypto/sha256"
	"errors"
	"fmt"
	"log"
)

func authByPassword(username string, password string) bool {
	user, err := getUser(username)

	if err != nil {
		log.Println(err)
		return false
	}

	hash := sha256.Sum256([]byte(password))
	hashString := fmt.Sprintf("%x", hash[:])
	return user.passwordHash == hashString
}

func checkPermission(user user, channel channel) error {
	for _, x := range channel.members {
		if x.name == user.name {
			return nil
		}
	}

	return errors.New("User is not a member of the channel")
}
