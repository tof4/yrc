package server

import (
	"fmt"
	"log"
	"net"
)

type clientNetworkInterface interface {
	sendData(data string)
	getAddress() net.Addr
}

type yrcClient struct {
	networkInterface clientNetworkInterface
	id               int
	nickname         string
}

func handleConnect(sender *yrcClient) {
	log.Println("Connected:", sender.networkInterface.getAddress())
	broadcast(sender, fmt.Sprintf("joined %d as %s", sender.id, sender.nickname))
}

func handleDisconnect(sender *yrcClient) {
	log.Println("Disconnected:", sender.networkInterface.getAddress())

	for i, c := range clients {
		if c.id == sender.id {
			clients[i] = clients[len(clients)-1]
			clients = clients[:len(clients)-1]
			break
		}
	}
}
