package main

import (
	"bufio"
	"net"
	"strings"

	"github.com/gotk3/gotk3/glib"
)

var connection net.Conn

func connect(address string) {
	tcpAddr, err := net.ResolveTCPAddr("tcp", address)
	if err != nil {
		writeToChat(err.Error())
		return
	}

	connection, err = net.DialTCP("tcp", nil, tcpAddr)
	if err != nil {
		writeToChat(err.Error())
		return
	}

	writeToChat("Connected with: " + connection.RemoteAddr().String())

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
