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

func sendToChannel(sender yrcClient, groupName string, content string) error {
	group, err := getGroup(groupName)

	if err != nil {
		replyWithError(sender, err)
		return err
	}

	formattedMessage := formatMessage(sender.username, groupName, content)

	for _, m := range group.members {
		receiver, err := getConnectedClientByUsername(m.name)
		if err == nil && sender.username != receiver.username {
			receiver.networkInterface.sendData(formattedMessage)
		}
	}

	saveMessage(groupName, formattedMessage)
	return nil
}

func getConnectedClientByUsername(username string) (yrcClient, error) {
	for _, c := range clients {
		if c.username == username {
			return c, nil
		}
	}
	return yrcClient{}, errors.New("User not found")
}

func formatMessage(groupName string, senderName string, content string) string {
	return fmt.Sprintf("%s:%d:%s:%s\n", groupName, time.Now().Unix(), senderName, content)
}
