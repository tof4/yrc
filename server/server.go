package server

import (
	"log"

	"github.com/gliderlabs/ssh"
)

var clients []yrcClient

func Initialize() {
	listenSsh()
}

func listenSsh() {
	ssh.Handle(func(s ssh.Session) {
		handleSshConnect(s)
	})

	log.Println("Starting SSH server on port 9999")
	log.Fatal(ssh.ListenAndServe(":9999", nil))
}
