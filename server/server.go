package server

import (
	"fmt"
	"log"

	"github.com/gliderlabs/ssh"
)

func Initialize() {
	openDatabase()
	listenSsh()
}

func listenSsh() {
	log.Println(fmt.Sprintf("Starting ssh server on port 9999"))

	ssh.Handle(func(s ssh.Session) {
		handleSshConnect(s)
	})

	ssh.ListenAndServe(":9999", nil,
		ssh.PasswordAuth(func(ctx ssh.Context, password string) bool {
			return authByPassword(ctx.User(), password)
		}),
	)
}
