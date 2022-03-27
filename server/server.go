package server

import (
	"fmt"
	"log"

	"github.com/gliderlabs/ssh"
)

var clients []yrcClient

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
		ssh.PasswordAuth(func(ctx ssh.Context, pass string) bool {
			user, err := getUserByUsername(ctx.User())
			if err != nil {
				log.Println(err)
				return false
			}
			return user.password == pass
		}),
	)
}
