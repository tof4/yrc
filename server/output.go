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

func sendToUser(user user, data string) {
	for _, x := range user.clients {
		x.sendData(data)
	}
}

func sendToChannel(senderClient client, channelName string, content string) error {
	channel, err := getChannel(channelName)
	sender, err := getClientUser(senderClient)

	if err != nil {
		replyWithError(senderClient, err)
		return err
	}

	formattedMessage := fmt.Sprintf("message %s %s %d %s\n", channel.name, sender.name, time.Now().Unix(), content)

	for _, m := range channel.members {
		sendToUser(m, formattedMessage)
	}

	saveMessage(channel.name, formattedMessage)
	return nil
}
