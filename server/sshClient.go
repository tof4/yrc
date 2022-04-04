package server

import (
	"bufio"
	"errors"
	"log"
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
		username:         session.User()}

	clients = append(clients, client)
	handleConnect(client)

	terminal := term.NewTerminal(session, "")
	for {
		line, err := terminal.ReadLine()
		if err != nil {

			if errors.As(err, &bufio.ErrFinalToken) {
				handleDisconnect(client)
				break
			} else {
				log.Println(err)
			}
			return
		}

		handleInput(line, client)
	}

}

func (client sshClient) sendData(data string) {
	client.session.Write([]byte(data))
}

func (client sshClient) getAddress() net.Addr {
	return client.session.RemoteAddr()
}

func (client sshClient) disconnect() {
	client.session.Close()
}
