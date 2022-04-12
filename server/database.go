package server

import (
	"bufio"
	"errors"
	"fmt"
	"os"
	"path/filepath"
	"strings"
	"time"
)

type databasePaths struct {
	root     string
	etc      string
	channels string
	passwd   string
}

type userModel struct {
	username     string
	passwordHash string
}

var (
	paths databasePaths
	users []userModel
)

func openDatabase(rootPath string) {
	paths.root = rootPath
	paths.etc = filepath.Join(paths.root, "etc")
	paths.channels = filepath.Join(paths.root, "chl")
	paths.passwd = filepath.Join(paths.etc, "passwd")

	err := os.MkdirAll(paths.etc, os.ModePerm)
	err = os.MkdirAll(paths.channels, os.ModePerm)
	_, err = os.OpenFile(paths.passwd, os.O_RDWR|os.O_CREATE, 0600)

	catchFatal(err)

	users = loadUsers()
}

func loadUsers() (newUsersList []userModel) {
	passwdFile, err := os.Open(paths.passwd)
	defer passwdFile.Close()
	catchFatal(err)
	scanner := bufio.NewScanner(passwdFile)
	scanner.Split(bufio.ScanLines)

	for scanner.Scan() {
		userProperties := strings.Split(scanner.Text(), ":")
		user := userModel{
			username:     userProperties[0],
			passwordHash: userProperties[1],
		}
		newUsersList = append(newUsersList, user)
	}

	return
}

func getUserPasswordHash(username string) (string, error) {
	for _, x := range users {
		if x.username == username {
			return x.passwordHash, nil
		}
	}

	return "", errors.New("User not found")
}

func getChannelMembers(channelName string) ([]string, error) {
	membersFile, err := os.ReadFile(filepath.Join(paths.channels, channelName, "members"))

	if err != nil && errors.Is(err, os.ErrNotExist) {
		return nil, errors.New("Channel not found")
	}
	if err != nil {
		catchFatal(err)
		return nil, nil
	}

	return strings.Split(string(membersFile), "\n"), nil
}

func saveMessage(channelName string, senderName string, content string) {
	path := filepath.Join(paths.channels, channelName, "chat")
	timestamp := time.Now().Format("2006-01-02 15:04:05")
	file, err := os.OpenFile(path, os.O_APPEND|os.O_WRONLY|os.O_CREATE, 0600)
	defer file.Close()
	_, err = file.WriteString(fmt.Sprintf("%s %s %s\n", timestamp, senderName, content))
	catchFatal(err)
}
