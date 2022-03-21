package server

import (
	"net"

	"github.com/gliderlabs/ssh"
	"golang.org/x/term"
)

type sshClient struct {
	session ssh.Session
}

func handleSshConnect(session ssh.Session) {

	newSshClient := sshClient{session: session}
	client := yrcClient{
		networkInterface: newSshClient,
		id:               len(clients),
		nickname:         "default"}

	clients = append(clients, client)

	terminal := term.NewTerminal(session, "")
	line := ""
	for {
		line, _ = terminal.ReadLine()
		handleCommand(line, &client)
	}
}

func (client sshClient) sendData(data string) {
	client.session.Write([]byte(data))
}

func (client sshClient) getAddress() net.Addr {
	return client.session.RemoteAddr()
}
