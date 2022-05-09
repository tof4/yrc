package database

import (
	"errors"
	"fmt"
	"os"
	"strings"

	"github.com/tof4/yrc/internal/errutil"
	"github.com/tof4/yrc/internal/strutil"
	"github.com/tof4/yrc/internal/validator"
	"golang.org/x/exp/slices"
)

func GetUser(name string) (*User, error) {
	for i, x := range Users {
		if x.Name == name {
			return &Users[i], nil
		}
	}

	return &User{}, errors.New("User not found")
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

func DeleteUser(username string) error {
	username = strutil.RemoveSpaces(username)
	_, err := GetUser(username)

	if err != nil {
		return errors.New("User doesn't exists")
	}

	for _, x := range Channels {
		RemoveFromChannel(x.Name, username)
	}

	index := slices.IndexFunc(Users, func(u User) bool { return u.Name == username })
	Users = append(Users[:index], Users[index+1:]...)
	refreshUsersFile()
	return nil
}

func refreshUsersFile() {
	var usersString string
	for _, x := range Users {
		var sb strings.Builder
		sb.WriteString(x.Name)
		sb.WriteString(" ")
		sb.WriteString(x.PasswordHash)
		usersString += fmt.Sprintln(sb.String())
	}

	usersFile, err := os.Create(Paths.Users)
	defer usersFile.Close()
	errutil.CatchFatal(err)
	usersFile.WriteString(usersString)
}
