package client

import (
	"fmt"
	"strconv"
	"strings"
)

func handleEvent(data string) {
	argumets := strings.Split(data, "/")

	switch argumets[0] {
	case "message":
		id, _ := strconv.Atoi(argumets[1])
		for i, c := range connectedUsers {
			if c.id == id {
				user := connectedUsers[i]
				logWriteMessage(fmt.Sprintf("%s: %s", user.nickname, strings.TrimSpace(argumets[2])))
				break
			}
		}

	case "join":
		var newUser yrcUser
		id, _ := strconv.Atoi(argumets[1])
		newUser.id = id
		newUser.nickname = strings.TrimSpace(argumets[2])
		connectedUsers = append(connectedUsers, newUser)
		logWriteStatus(fmt.Sprintf("%s joined", newUser.nickname))
	}
}
