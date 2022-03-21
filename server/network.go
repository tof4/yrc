package server

import "fmt"

func broadcast(sender *yrcClient, data string) {

	receivers := clients

	for i, c := range receivers {
		if c.id == sender.id {
			receivers[i] = receivers[len(receivers)-1]
			receivers = receivers[:len(receivers)-1]
			break
		}
	}

	for _, receiver := range receivers {
		receiver.networkInterface.sendData(fmt.Sprintf("%s\n", data))
	}
}
