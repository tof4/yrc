package client

import (
	"bufio"
	"net"
)

type yrcUser struct {
	id       int
	nickname string
}

var (
	connection     net.Conn
	connectedUsers []yrcUser
)

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
			handleEvent(str)
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
