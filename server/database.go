package server

import (
	"errors"
	"fmt"
	"os"
	"path/filepath"
	"strings"
	"time"
)

type databasePaths struct {
	root     string
	users    string
	channels string
}

var paths databasePaths

func openDatabase() {
	paths.root = "ydb"
	paths.users = filepath.Join(paths.root, "usr")
	paths.channels = filepath.Join(paths.root, "chl")
	err := os.MkdirAll(paths.users, os.ModePerm)
	err = os.MkdirAll(paths.channels, os.ModePerm)
	catchFatal(err)
}

func getUserPasswordHash(username string) (string, error) {
	usersDir, err := os.Open(paths.users)
	catchFatal(err)
	defer usersDir.Close()

	list, _ := usersDir.Readdirnames(0)
	for _, name := range list {
		if name == username {
			passwordHash, err := os.ReadFile(filepath.Join(paths.users, name, "passwordhash"))
			catchFatal(err)
			return strings.TrimSpace(string(passwordHash)), nil
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
