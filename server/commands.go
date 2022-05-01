package server

import (
	"strconv"
	"strings"
)

func callCommand(client client, argumets []string) {
	switch argumets[0] {
	case "send":
		send(client, argumets)

	case "exit":
		exit(client)

	case "read":
		read(client, argumets)
	}
}

func send(client client, argumets []string) {
	if len(argumets) < 3 {
		return
	}

	sendToChannel(client, argumets[1], argumets[2])
}

func exit(client client) {
	client.disconnect()
}

func read(client client, argumets []string) {
	if len(argumets) < 3 {
		return
	}

	channelName := argumets[1]
	amount, err := strconv.Atoi(argumets[2])

	if err != nil {
		replyWithError(client, err)
		return
	}

	messages, err := getChannelMessages(channelName, amount)

	if err != nil {
		replyWithError(client, err)
		return
	}

	sendToClient(client, strings.Join(messages, "\n"))
}
