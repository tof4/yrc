package database

import (
	"bufio"
	"errors"
	"log"
	"os"
	"path/filepath"
	"strings"

	"github.com/tof4/yrc/pkg/common"
)

func OpenDatabase(rootPath string) {
	Paths.Root = rootPath
	Paths.Etc = filepath.Join(Paths.Root, "etc")
	Paths.Chat = filepath.Join(Paths.Root, "chat")
	Paths.Users = filepath.Join(Paths.Etc, "users")
	Paths.Channels = filepath.Join(Paths.Etc, "channels")
	Paths.Key = filepath.Join(Paths.Etc, "key.pem")

	err := os.MkdirAll(Paths.Etc, os.ModePerm)
	err = os.MkdirAll(Paths.Chat, os.ModePerm)
	_, err = os.OpenFile(Paths.Users, os.O_RDWR|os.O_CREATE, 0600)
	_, err = os.OpenFile(Paths.Channels, os.O_RDWR|os.O_CREATE, 0600)

	common.CatchFatal(err)

	Users = loadUsers()
	channels = loadChannels(Users)

	log.Printf("Loaded %d user(s) and %d channel(s)\n", len(Users), len(channels))
}

func GetUser(name string) (User, error) {
	for i, x := range Users {
		if x.Name == name {
			return Users[i], nil
		}
	}

	return User{}, errors.New("User not found")
}

func GetChannel(name string) (Channel, error) {
	for _, x := range channels {
		if x.Name == name {
			return x, nil
		}
	}

	return Channel{}, errors.New("Channel not found")
}

func SaveMessage(channelName string, message string) {
	path := filepath.Join(Paths.Chat, channelName)
	file, err := os.OpenFile(path, os.O_APPEND|os.O_WRONLY|os.O_CREATE, 0600)
	defer file.Close()
	_, err = file.WriteString(message)
	common.CatchFatal(err)
}

func GetChannelMessages(channelName string, amount int) ([]string, error) {
	_, err := GetChannel(channelName)

	if err != nil {
		return []string{}, err
	}

	if amount < 1 || amount > 1000 {
		return []string{}, errors.New("Invalid amount")
	}
	amount++

	path := filepath.Join(Paths.Chat, channelName)
	return BackwardFileRead(path, amount), nil
}

func loadUsers() (newUsersList []User) {
	usersFile, err := os.Open(Paths.Users)
	defer usersFile.Close()
	common.CatchFatal(err)

	scanner := bufio.NewScanner(usersFile)
	scanner.Split(bufio.ScanLines)

	for scanner.Scan() {
		userProperties := strings.Split(scanner.Text(), " ")
		user := User{
			Name:         userProperties[0],
			PasswordHash: userProperties[1],
		}
		newUsersList = append(newUsersList, user)
	}

	return
}

func loadChannels(users []User) (channelsList []Channel) {
	channelsFile, err := os.Open(Paths.Channels)
	defer channelsFile.Close()
	common.CatchFatal(err)

	scanner := bufio.NewScanner(channelsFile)
	scanner.Split(bufio.ScanLines)

	for scanner.Scan() {
		channelProperties := strings.Split(scanner.Text(), " ")
		channelMembersStrings := strings.Split(channelProperties[1], ",")

		channel := Channel{
			Name:    channelProperties[0],
			Members: []*User{},
		}

		for _, x := range channelMembersStrings {
			user, err := GetUser(x)
			if err == nil {
				channel.Members = append(channel.Members, &user)
			}
		}

		channelsList = append(channelsList, channel)
	}

	return
}
