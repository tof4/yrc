package server

import (
	"crypto/sha256"
	"errors"
	"fmt"
	"log"

	"github.com/tof4/yrc/pkg/database"
)

func authByPassword(username string, password string) bool {
	user, err := database.GetUser(username)

	if err != nil {
		log.Println(err)
		return false
	}

	hash := sha256.Sum256([]byte(password))
	hashString := fmt.Sprintf("%x", hash[:])
	return user.PasswordHash == hashString
}

func checkPermission(username string, channel database.Channel) error {
	for _, x := range channel.Members {
		if x.Name == username {
			return nil
		}
	}

	return errors.New("User is not a member of the channel")
}
