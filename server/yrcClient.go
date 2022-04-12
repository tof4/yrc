package server

import (
	"fmt"
	"log"
	"net"
)

type clientNetworkInterface interface {
	sendData(data string)
	getAddress() net.Addr
	disconnect()
}

type yrcClient struct {
	networkInterface clientNetworkInterface
	username         string
}

func handleConnect(client yrcClient) {
	log.Println(fmt.Sprintf("Connected %s from %s", client.username, client.networkInterface.getAddress()))
}

func handleDisconnect(client yrcClient) {
	log.Println(fmt.Sprintf("Disconnected %s", client.username))

	for i, c := range clients {
		if c.username == client.username {
			clients[i] = clients[len(clients)-1]
			clients = clients[:len(clients)-1]
			return
		}
	}
}
