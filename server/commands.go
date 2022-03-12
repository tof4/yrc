package main

import (
	"strings"
)

func handleCommand(command string, sender *yrcClient) {
	argumets := strings.Split(command, "/")

	switch argumets[0] {
	case "send":
		for _, client := range server.clients {
			if client.id != sender.id {
				client.connection.Write([]byte(sender.nickname + ": " + argumets[1]))
			}
		}
	case "nick":
		sender.nickname = strings.TrimSpace(argumets[1])
	}
}
