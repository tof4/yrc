package server

import (
	"errors"
	"fmt"
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

func sendToChannel(sender yrcClient, channelName string, content string) error {
	members, err := getChannelMembers(channelName)
	saveMessage(channelName, sender.username, content)
	for _, m := range members {
		receiver, err := getConnectedClientByUsername(m)
		if err == nil && sender.username != receiver.username {
			receiver.networkInterface.sendData(content)
		}
	}
	catchFatal(err)
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
