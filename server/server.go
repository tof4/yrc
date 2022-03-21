package server

import (
	"log"
	"net"

	"github.com/gliderlabs/ssh"
)

var tcpListener net.Listener
var clients []yrcClient

func Initialize() {
	go listenSsh()
	startServer()
	listenTcp()
	defer tcpListener.Close()
}

func startServer() {
	log.Println("Starting TCP server on port 9999")
	l, err := net.Listen("tcp", ":9999")
	if err != nil {
		log.Fatal(err)
		return
	}

	tcpListener = l
}

func listenTcp() {
	for {
		c, err := tcpListener.Accept()
		if err != nil {
			log.Fatal(err)
			return
		}

		go handleTcpConnect(c)
	}
}

func listenSsh() {
	ssh.Handle(func(s ssh.Session) {
		handleSshConnect(s)
	})

	log.Println("Starting SSH server on port 9998")
	log.Fatal(ssh.ListenAndServe(":9998", nil))
}
