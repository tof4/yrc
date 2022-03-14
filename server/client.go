package server

import (
	"bufio"
	"errors"
	"log"
	"net"
)

type yrcClient struct {
	connection net.Conn
	id         int
	nickname   string
}

func handleConnect(connection net.Conn) {
	log.Println("New connection:", connection.RemoteAddr())

	client := yrcClient{
		connection: connection,
		id:         len(clients),
		nickname:   "default"}

	clients = append(clients, client)
	reader := bufio.NewReader(connection)
	eventJoin(&client)

	for {
		command, err := reader.ReadString('\n')
		if err != nil {

			if errors.As(err, &bufio.ErrFinalToken) {
				handleDisconnect(client)
				break
			} else {
				log.Println(err)
			}
			return
		}

		handleCommand(string(command), &client)
	}
}

func handleDisconnect(client yrcClient) {
	log.Println("Disconnected:", client.connection.RemoteAddr())

	for i, c := range clients {
		if c.id == client.id {
			clients[i] = clients[len(clients)-1]
			clients = clients[:len(clients)-1]
			break
		}
	}
}
