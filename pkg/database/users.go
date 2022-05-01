package database

import (
	"errors"
	"fmt"
	"os"
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

	if len(username) < 1 || len(username) > 20 {
		return errors.New("Invalid username length")
	}

	if len(password) < 1 || len(password) > 100 {
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

	file, err := os.OpenFile(Paths.Users, os.O_APPEND|os.O_WRONLY|os.O_CREATE, 0600)
	defer file.Close()
	_, err = file.WriteString(userString)
	common.CatchFatal(err)

	return nil
}
