package database

import (
	"errors"
	"fmt"

	"github.com/tof4/yrc/internal/strutil"
	"github.com/tof4/yrc/internal/validator"
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
	username = strutil.RemoveSpaces(username)

	if validator.ValidateLength(username, 1, 20) {
		return errors.New("Invalid username length")
	}

	if validator.ValidateLength(password, 1, 100) {
		return errors.New("Invalid password length")
	}

	_, err := GetUser(username)

	if err == nil {
		return errors.New("Username already exists")
	}

	passwordHash := strutil.Sha256(password)
	userString := fmt.Sprintf("%s %s\n", username, passwordHash)

	Users = append(Users,
		User{
			Name:         username,
			PasswordHash: passwordHash,
		})

	fileAppend(userString, Paths.Users)
	return nil
}
