package main

import (
	"bufio"
	"log"
	"net"
	"strings"

	"github.com/gotk3/gotk3/glib"
)

var connection net.Conn

func connect(address string) {
	tcpAddr, err := net.ResolveTCPAddr("tcp", address)
	if err != nil {
		log.Println("Resolve failed:", err.Error())
	}

	connection, err = net.DialTCP("tcp", nil, tcpAddr)
	if err != nil {
		log.Println("Dial failed:", err.Error())
	}

	connbuf := bufio.NewReader(connection)
	for {
		str, err := connbuf.ReadString('\n')
		if err != nil {
			break
		}

		if len(str) > 0 {
			glib.IdleAdd(func() {
				writeToChat(strings.TrimSpace(str))
			})
		}
	}
}

func sendMessage(message string) {
	connection.Write([]byte("send/" + message + "\n"))
}
