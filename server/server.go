package server

import (
	"log"
	"net"
)

var listener net.Listener
var clients []yrcClient

func Initialize() {
	log.Println("Starting YRC")
	startServer()
	listenClients()
	defer listener.Close()
}

func startServer() {
	l, err := net.Listen("tcp", ":9999")
	if err != nil {
		log.Fatal(err)
		return
	}

	listener = l
}

func listenClients() {
	for {
		connection, err := listener.Accept()
		if err != nil {
			log.Fatal(err)
			return
		}

		go handleConnect(connection)
	}
}
