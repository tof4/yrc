package main

import (
	"bufio"
	"errors"
	"log"
	"net"
)

type yrcServer struct {
	listener net.Listener
	clients  []yrcClient
}

type yrcClient struct {
	connection net.Conn
	id         int
	nickname   string
}

var server yrcServer

func main() {
	log.Println("Starting YRC")
	startServer()
	listenClients()
	defer server.listener.Close()
}

func startServer() {
	l, err := net.Listen("tcp", ":9999")
	if err != nil {
		log.Fatal(err)
		return
	}

	server = yrcServer{listener: l}
}

func listenClients() {
	for {
		connection, err := server.listener.Accept()
		if err != nil {
			log.Fatal(err)
			return
		}

		go handleConnection(connection)
	}
}

func handleConnection(connection net.Conn) {
	log.Println("New connection:", connection.RemoteAddr())

	client := yrcClient{
		connection: connection,
		id:         len(server.clients),
		nickname:   "default"}

	server.clients = append(server.clients, client)
	reader := bufio.NewReader(connection)

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
