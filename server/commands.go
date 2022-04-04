package server

func handleCommand(command string, client yrcClient) {
	argumets := getStringsBetweenQuotes(command)

	if len(argumets) < 1 {
		return
	}

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

	if validateMessage(argumets[2]) {
		sendToChannel(client, argumets[1], argumets[2])
	}
}

func exit(client yrcClient) {
	client.networkInterface.disconnect()
}
