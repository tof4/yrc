package database

import (
	"bufio"
	"os"
	"path/filepath"
	"strings"

	"github.com/tof4/yrc/internal/errutil"
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

	errutil.CatchFatal(err)

	Users = loadUsers()
	Channels = loadChannels(Users)
}

func loadUsers() (newUsersList []User) {
	usersFile, err := os.Open(Paths.Users)
	defer usersFile.Close()
	errutil.CatchFatal(err)

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
	errutil.CatchFatal(err)

	scanner := bufio.NewScanner(channelsFile)
	scanner.Split(bufio.ScanLines)

	for scanner.Scan() {
		channelProperties := strings.Split(scanner.Text(), " ")
		channel := Channel{
			Name: channelProperties[0],
		}

		if len(channelProperties) != 1 {
			channelMembersString := strings.Split(channelProperties[1], ",")
			for _, x := range channelMembersString {
				user, err := GetUser(x)
				if err == nil {
					channel.Members = append(channel.Members, user)
				}
			}
		}

		channelsList = append(channelsList, channel)
	}
	return
}
