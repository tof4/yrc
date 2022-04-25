package server

import (
	"net"

	"github.com/gliderlabs/ssh"
	"github.com/google/uuid"
)

type client struct {
	session ssh.Session
	id      uuid.UUID
}

func (client client) sendData(data string) {
	client.session.Write([]byte(data))
}

func (client client) getAddress() net.Addr {
	return client.session.RemoteAddr()
}

func (client client) disconnect() {
	client.session.Close()
}
