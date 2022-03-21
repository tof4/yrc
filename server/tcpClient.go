package server

import (
	"bufio"
	"errors"
	"log"
	"net"
)

type tcpClient struct {
	session net.Conn
}

func handleTcpConnect(connection net.Conn) {

	newTcpClient := tcpClient{
		session: connection,
	}
	client := yrcClient{
		networkInterface: newTcpClient,
		id:               len(clients),
		nickname:         "default"}

	clients = append(clients, client)
	handleConnect(client)

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

		handleCommand(command, client)
	}
}

func (client tcpClient) sendData(data string) {
	client.session.Write([]byte(data))
}

func (client tcpClient) getAddress() net.Addr {
	return client.session.RemoteAddr()
}

func (client tcpClient) disconnect() {
	client.session.Close()
}
