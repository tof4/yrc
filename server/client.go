package main

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
		id:         len(server.clients),
		nickname:   "default"}

	server.clients = append(server.clients, client)
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

	for i, c := range server.clients {
		if c.id == client.id {
			server.clients[i] = server.clients[len(server.clients)-1]
			server.clients = server.clients[:len(server.clients)-1]
			break
		}
	}
}
