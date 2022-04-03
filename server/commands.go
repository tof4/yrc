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

	case "exit":
		exit(client)
	}
}

func send(client yrcClient, argumets []string) {
	message := getStringBetweenQuotes(strings.Join(argumets, " "))
	if validateMessage(message) {
		timestamp := time.Now().Format("2006-01-02|15:04:05")
		broadcast(client, fmt.Sprintf(`message from %s at %s "%s"`, client.username, timestamp, message))
	}
}

func exit(client yrcClient) {
	client.networkInterface.disconnect()
}
