package server

import "fmt"

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
