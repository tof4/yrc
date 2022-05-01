package server

import (
	"fmt"
	"time"

	"github.com/tof4/yrc/pkg/database"
)

func sendToAll(sender client, data string) {
	for _, x := range clients {
		if x.id != sender.id {
			x.sendData(fmt.Sprintf("%s\n", data))
		}
	}
}

func sendToUser(username string, ignoredClient client, data string) {
	for _, x := range clients {
		if x.username == username && x.id != ignoredClient.id {
			x.sendData(data)
		}
	}
}

func sendToClient(client client, data string) {
	client.sendData(data)
}

func sendToChannel(sender client, channelName string, content string) {
	channel, err := database.GetChannel(channelName)

	if err != nil {
		replyWithError(sender, err)
		return
	}

	err = checkPermission(sender.username, channel)

	if err != nil {
		replyWithError(sender, err)
		return
	}

	formattedMessage := fmt.Sprintf("message %s %s %d %s\n", channel.Name, sender.username, time.Now().Unix(), content)

	for _, m := range channel.Members {
		sendToUser(m.Name, sender, formattedMessage)
	}

	database.SaveMessage(channel.Name, formattedMessage)
}
