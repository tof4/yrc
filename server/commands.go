package server

import (
	"fmt"
	"strings"
	"time"

	"github.com/tof4/yrc/common"
)

func handleCommand(command string, sender *yrcClient) {
	argumets := strings.Split(command, " ")

	switch argumets[0] {
	case "send":
		send(sender, argumets)

	case "nick":
		nick(sender, argumets)
	}
}

func send(sender *yrcClient, argumets []string) {
	message := common.GetStringBetween(strings.Join(argumets, " "), "'")
	timestamp := time.Now().Format("2006-01-02 15:04:05")
	broadcast(*sender, fmt.Sprintf("message from %d at %s %s", sender.id, timestamp, message))
}

func nick(sender *yrcClient, argumets []string) {
	sender.nickname = strings.TrimSpace(argumets[1])
}
