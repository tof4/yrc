package server

func callCommand(client yrcClient, argumets []string) {
	switch argumets[0] {
	case "send":
		send(client, argumets)

	case "exit":
		exit(client)
	}
}

func send(client yrcClient, argumets []string) {
	if len(argumets) < 3 {
		return
	}

	sendToChannel(client, argumets[1], argumets[2])
}

func exit(client yrcClient) {
	client.networkInterface.disconnect()
}
