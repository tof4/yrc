package main

import "fmt"

func broadcast(sender yrcClient, data string) {

	receivers := server.clients

	for i, c := range receivers {
		if c.id == sender.id {
			receivers[i] = receivers[len(receivers)-1]
			receivers = receivers[:len(receivers)-1]
			break
		}
	}

	for _, receiver := range receivers {
		receiver.connection.Write([]byte(fmt.Sprintf("%s\n", data)))
	}
}
