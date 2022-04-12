package server

import (
	"bufio"
	"errors"
	"fmt"
	"log"
	"os"
	"path/filepath"
	"strings"
	"time"
)

type databasePaths struct {
	root   string
	etc    string
	chat   string
	passwd string
	group  string
}

type user struct {
	name         string
	passwordHash string
}

type group struct {
	name    string
	members []user
}

var (
	paths  databasePaths
	users  []user
	groups []group
)

func openDatabase(rootPath string) {
	paths.root = rootPath
	paths.etc = filepath.Join(paths.root, "etc")
	paths.chat = filepath.Join(paths.root, "chat")
	paths.passwd = filepath.Join(paths.etc, "passwd")
	paths.group = filepath.Join(paths.etc, "group")

	err := os.MkdirAll(paths.etc, os.ModePerm)
	err = os.MkdirAll(paths.chat, os.ModePerm)
	_, err = os.OpenFile(paths.passwd, os.O_RDWR|os.O_CREATE, 0600)
	_, err = os.OpenFile(paths.group, os.O_RDWR|os.O_CREATE, 0600)

	catchFatal(err)

	users = loadUsers()
	groups = loadGroups(users)

	log.Printf("Loaded %d user(s) and %d group(s)\n", len(users), len(groups))
}

func loadUsers() (newUsersList []user) {
	passwdFile, err := os.Open(paths.passwd)
	defer passwdFile.Close()
	catchFatal(err)

	scanner := bufio.NewScanner(passwdFile)
	scanner.Split(bufio.ScanLines)

	for scanner.Scan() {
		userProperties := strings.Split(scanner.Text(), ":")
		user := user{
			name:         userProperties[0],
			passwordHash: userProperties[1],
		}
		newUsersList = append(newUsersList, user)
	}

	return
}

func loadGroups(users []user) (newGroupsList []group) {
	groupsFile, err := os.Open(paths.group)
	defer groupsFile.Close()
	catchFatal(err)

	scanner := bufio.NewScanner(groupsFile)
	scanner.Split(bufio.ScanLines)

	for scanner.Scan() {
		groupProperties := strings.Split(scanner.Text(), ":")
		groupMembersStrings := strings.Split(groupProperties[1], ",")

		group := group{
			name:    groupProperties[0],
			members: []user{},
		}

		for _, x := range groupMembersStrings {
			user, err := getUser(x)
			if err == nil {
				group.members = append(group.members, user)
			}
		}

		newGroupsList = append(newGroupsList, group)
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

func getGroup(name string) (group, error) {
	for _, x := range groups {
		if x.name == name {
			return x, nil
		}
	}

	return group{}, errors.New("Group not found")
}

func saveMessage(channelName string, senderName string, content string) {
	path := filepath.Join(paths.chat, channelName)
	timestamp := time.Now().Format("2006-01-02 15:04:05")
	file, err := os.OpenFile(path, os.O_APPEND|os.O_WRONLY|os.O_CREATE, 0600)
	defer file.Close()
	_, err = file.WriteString(fmt.Sprintf("%s %s: %s\n", timestamp, senderName, content))
	catchFatal(err)
}
