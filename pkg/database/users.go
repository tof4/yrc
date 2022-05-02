package database

import (
	"errors"
	"fmt"
	"strings"

	"github.com/tof4/yrc/internal/common"
)

func GetUser(name string) (User, error) {
	for i, x := range Users {
		if x.Name == name {
			return Users[i], nil
		}
	}

	return User{}, errors.New("User not found")
}

func CreateUser(username string, password string) error {
	username = strings.TrimSpace(username)
	username = strings.Split(username, " ")[0]

	if common.ValidateLength(username, 1, 20) {
		return errors.New("Invalid username length")
	}

	if common.ValidateLength(password, 1, 100) {
		return errors.New("Invalid password length")
	}

	_, err := GetUser(username)

	if err == nil {
		return errors.New("Username already in use")
	}

	passwordHash := common.Sha256String(password)
	userString := fmt.Sprintf("%s %s\n", username, passwordHash)

	Users = append(Users,
		User{
			Name:         username,
			PasswordHash: passwordHash,
		})

	fileAppend(userString, Paths.Users)
	return nil
}
