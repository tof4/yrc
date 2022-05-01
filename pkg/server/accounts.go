package server

import (
	"errors"
	"log"

	"github.com/tof4/yrc/internal/common"
	"github.com/tof4/yrc/pkg/database"
)

func authByPassword(username string, password string) bool {
	user, err := database.GetUser(username)

	if err != nil {
		log.Println(err)
		return false
	}

	return user.PasswordHash == common.Sha256String(password)
}

func checkPermission(username string, channel database.Channel) error {
	for _, x := range channel.Members {
		if x.Name == username {
			return nil
		}
	}

	return errors.New("User is not a member of the channel")
}
