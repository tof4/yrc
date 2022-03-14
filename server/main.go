package main

import (
	"log"
	"net"
)

type yrcServer struct {
	listener net.Listener
	clients  []yrcClient
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

		go handleConnect(connection)
	}
}
