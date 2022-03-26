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
	id               int
	nickname         string
}

func handleConnect(client yrcClient) {
	log.Println("Connected:", client.networkInterface.getAddress())
	broadcast(client, fmt.Sprintf(`connected %d as "%s"`, client.id, client.nickname))
}

func handleDisconnect(client yrcClient) {
	log.Println("Disconnected:", client.networkInterface.getAddress())
	broadcast(client, fmt.Sprintf(`disconnected "%d"`, client.id))

	for i, c := range clients {
		if c.id == client.id {
			clients[i] = clients[len(clients)-1]
			clients = clients[:len(clients)-1]
			break
		}
	}
}
