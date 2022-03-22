package server

import (
	"fmt"
	"strings"
	"time"
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
	message := getStringBetweenQuotes(strings.Join(argumets, " "))
	timestamp := time.Now().Format("2006-01-02 15:04:05")
	broadcast(client, fmt.Sprintf("message from %d at %s '%s'", client.id, timestamp, message))
}

func nick(client yrcClient, argumets []string) {
	oldNickname := client.nickname
	client.nickname = getStringBetweenQuotes(strings.Join(argumets, " "))
	broadcast(client, fmt.Sprintf("renamed %d from %s to '%s'", client.id, oldNickname, client.nickname))
}

func exit(client yrcClient) {
	client.networkInterface.disconnect()
}
