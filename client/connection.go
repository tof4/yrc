package main

import (
	"bufio"
	"net"
	"strings"
)

var connection net.Conn

func connect(address string) {
	tcpAddr, err := net.ResolveTCPAddr("tcp", address)
	if err != nil {
		logWriteError("Connection error: " + err.Error())
		return
	}

	connection, err = net.DialTCP("tcp", nil, tcpAddr)
	if err != nil {
		logWriteError("Connection error: " + err.Error())
		return
	}

	switchConnectButton()
	logWriteStatus("Connected with: " + connection.RemoteAddr().String())

	connbuf := bufio.NewReader(connection)
	for {
		str, err := connbuf.ReadString('\n')
		if err != nil {
			break
		}

		if len(str) > 0 {
			logWriteMessage(strings.TrimSpace(str))
		}
	}

	switchConnectButton()
	logWriteStatus("Disconnected from: " + connection.RemoteAddr().String())
}

func sendMessage(message string) {
	if connection != nil {
		connection.Write([]byte("send/" + message + "\n"))
	}
}
