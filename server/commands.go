package server

import (
	"strings"
)

func handleCommand(command string, sender *yrcClient) {
	argumets := strings.Split(command, "/")

	switch argumets[0] {
	case "send":
		eventMessage(sender, argumets[1])
	case "nick":
		sender.nickname = strings.TrimSpace(argumets[1])
	}
}
