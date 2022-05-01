package server

import (
	"bufio"
	"errors"
	"log"
	"os"
	"path/filepath"
	"strings"
	"unicode/utf8"
)

type databasePaths struct {
	root     string
	etc      string
	chat     string
	users    string
	channels string
	key      string
}

type channel struct {
	name    string
	members []*user
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
	paths.key = filepath.Join(paths.etc, "key.pem")

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
			members: []*user{},
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

func getUser(name string) (*user, error) {
	for i, x := range users {
		if x.name == name {
			return &users[i], nil
		}
	}

	return &user{}, errors.New("User not found")
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

func getChannelMessages(channel string, count int) []string {
	path := filepath.Join(paths.chat, channel)
	return BackwardFileRead(path, count)
}

func BackwardFileRead(path string, count int) []string {
	file, _ := os.Open(path)
	defer file.Close()

	buf := make([]byte, 1)
	lines := make([]string, count)
	var sb strings.Builder
	start, _ := file.Seek(0, 2)

	currentLine := count - 1

	for i := start; i >= 0; i-- {
		if currentLine == -1 {
			break
		}

		file.ReadAt(buf, i)

		c, _ := utf8.DecodeRune(buf)
		if c == utf8.RuneError && len(buf) <= 5 {
			buf = make([]byte, len(buf)+1)
		} else {
			if c == '\n' {
				lines[currentLine] = reverse(sb.String())
				sb.Reset()
				buf = make([]byte, 1)
				currentLine--
			} else {
				sb.WriteRune(c)
			}
		}
	}

	return lines
}

func reverse(s string) string {
	runes := []rune(s)
	for i, j := 0, len(runes)-1; i < j; i, j = i+1, j-1 {
		runes[i], runes[j] = runes[j], runes[i]
	}
	return string(runes)
}
