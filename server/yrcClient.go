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
	log.Println("Connected:", client.networkInterface.getAddress())
	broadcast(client, fmt.Sprintf(`connected %s`, client.username))
}

func handleDisconnect(client yrcClient) {
	log.Println("Disconnected:", client.networkInterface.getAddress())
	broadcast(client, fmt.Sprintf(`disconnected %s`, client.username))

	for i, c := range clients {
		if c.username == client.username {
			clients[i] = clients[len(clients)-1]
			clients = clients[:len(clients)-1]
			return
		}
	}
}
