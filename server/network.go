package server

import (
	"errors"
	"fmt"
	"time"
)

var clients []yrcClient

func broadcast(sender yrcClient, data string) {

	receivers := clients

	for i, c := range receivers {
		if c.username == sender.username {
			receivers[i] = receivers[len(receivers)-1]
			receivers = receivers[:len(receivers)-1]
			return
		}
	}

	for _, receiver := range receivers {
		receiver.networkInterface.sendData(fmt.Sprintf("%s\n", data))
	}
}

func sendToChannel(sender yrcClient, channel string, content string) error {
	group, err := getChannel(channel)

	if err != nil {
		replyWithError(sender, err)
		return err
	}

	formattedMessage := fmt.Sprintf("message %s %s %d %s\n", group.name, sender.username, time.Now().Unix(), content)

	for _, m := range group.members {
		receiver, err := getClientByUsername(m.name)
		if err == nil && sender.username != receiver.username {
			receiver.networkInterface.sendData(formattedMessage)
		}
	}

	saveMessage(channel, formattedMessage)
	return nil
}

func getClientByUsername(username string) (yrcClient, error) {
	for _, c := range clients {
		if c.username == username {
			return c, nil
		}
	}
	return yrcClient{}, errors.New("User not found")
}
