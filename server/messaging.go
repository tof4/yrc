package main

func handleMessage(message string, sender yrcClient) {
	for _, client := range server.clients {
		if client.id != sender.id {
			client.connection.Write([]byte(message))
		}
	}
}
