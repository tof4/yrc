package server

import (
	"fmt"
	"time"
)

func sendToAll(sender client, data string) {
	for _, x := range users {
		for _, y := range x.clients {
			if y.id != sender.id {
				y.sendData(fmt.Sprintf("%s\n", data))
			}
		}
	}
}

func sendToUser(user *user, client client, data string) {
	for _, x := range user.clients {
		if x.id != client.id {
			x.sendData(data)
		}
	}
}

func sendToChannel(sender client, channelName string, content string) {
	channel, err := getChannel(channelName)

	if err != nil {
		replyWithError(sender, err)
		return
	}

	err = checkPermission(*sender.user, channel)

	if err != nil {
		replyWithError(sender, err)
		return
	}

	formattedMessage := fmt.Sprintf("message %s %s %d %s\n", channel.name, sender.user.name, time.Now().Unix(), content)

	for _, m := range channel.members {
		sendToUser(m, sender, formattedMessage)
	}

	saveMessage(channel.name, formattedMessage)
}
