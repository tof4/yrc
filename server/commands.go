package server

import (
	"fmt"
	"strings"
	"time"

	"github.com/tof4/yrc/common"
)

func handleCommand(command string, client yrcClient) {
	argumets := strings.Split(strings.TrimSpace(command), " ")

	switch argumets[0] {
	case "send":
		send(client, argumets)

	case "nick":
		nick(client, argumets)

	case "exit":
		exit(client)
	}
}

func send(client yrcClient, argumets []string) {
	message := common.GetStringBetween(strings.Join(argumets, " "), "'")
	timestamp := time.Now().Format("2006-01-02 15:04:05")
	broadcast(client, fmt.Sprintf("message from %d at %s %s", client.id, timestamp, message))
}

func nick(client yrcClient, argumets []string) {
	client.nickname = strings.TrimSpace(argumets[1])
}

func exit(client yrcClient) {
	client.networkInterface.disconnect()
}
