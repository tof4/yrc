package server

import (
	"bufio"
	"errors"
	"log"
	"os"
	"path/filepath"
	"strings"
)

type databasePaths struct {
	root     string
	etc      string
	chat     string
	users    string
	channels string
}

type channel struct {
	name    string
	members []user
}

var (
	paths    databasePaths
	users    []user
	channels []channel
)

func openDatabase(rootPath string) {
	paths.root = rootPath
	paths.etc = filepath.Join(paths.root, "etc")
	paths.chat = filepath.Join(paths.root, "chat")
	paths.users = filepath.Join(paths.etc, "users")
	paths.channels = filepath.Join(paths.etc, "channels")

	err := os.MkdirAll(paths.etc, os.ModePerm)
	err = os.MkdirAll(paths.chat, os.ModePerm)
	_, err = os.OpenFile(paths.users, os.O_RDWR|os.O_CREATE, 0600)
	_, err = os.OpenFile(paths.channels, os.O_RDWR|os.O_CREATE, 0600)

	catchFatal(err)

	users = loadUsers()
	channels = loadChannels(users)

	log.Printf("Loaded %d user(s) and %d channel(s)\n", len(users), len(channels))
}

func loadUsers() (newUsersList []user) {
	usersFile, err := os.Open(paths.users)
	defer usersFile.Close()
	catchFatal(err)

	scanner := bufio.NewScanner(usersFile)
	scanner.Split(bufio.ScanLines)

	for scanner.Scan() {
		userProperties := strings.Split(scanner.Text(), " ")
		user := user{
			name:         userProperties[0],
			passwordHash: userProperties[1],
		}
		newUsersList = append(newUsersList, user)
	}

	return
}

func loadChannels(users []user) (channelsList []channel) {
	channelsFile, err := os.Open(paths.channels)
	defer channelsFile.Close()
	catchFatal(err)

	scanner := bufio.NewScanner(channelsFile)
	scanner.Split(bufio.ScanLines)

	for scanner.Scan() {
		channelProperties := strings.Split(scanner.Text(), " ")
		channelMembersStrings := strings.Split(channelProperties[1], ",")

		channel := channel{
			name:    channelProperties[0],
			members: []user{},
		}

		for _, x := range channelMembersStrings {
			user, err := getUser(x)
			if err == nil {
				channel.members = append(channel.members, user)
			}
		}

		channelsList = append(channelsList, channel)
	}

	return
}

func getUser(name string) (user, error) {
	for _, x := range users {
		if x.name == name {
			return x, nil
		}
	}

	return user{}, errors.New("User not found")
}

func getChannel(name string) (channel, error) {
	for _, x := range channels {
		if x.name == name {
			return x, nil
		}
	}

	return channel{}, errors.New("Channel not found")
}

func saveMessage(channelName string, message string) {
	path := filepath.Join(paths.chat, channelName)
	file, err := os.OpenFile(path, os.O_APPEND|os.O_WRONLY|os.O_CREATE, 0600)
	defer file.Close()
	_, err = file.WriteString(message)
	catchFatal(err)
}
