package server

import (
	"fmt"
	"log"

	"github.com/gliderlabs/ssh"
)

func Initialize(port int, rootPath string) {
	log.Println(fmt.Sprintf("Starting ssh server on port %d", port))
	log.Println(fmt.Sprintf("Database set in %s", rootPath))
	openDatabase(rootPath)
	listenSsh(port)
}

func listenSsh(port int) {
	ssh.Handle(func(s ssh.Session) {
		handleSshConnect(s)
	})
	ssh.ListenAndServe(fmt.Sprintf(":%d", port), nil,
		ssh.PasswordAuth(func(ctx ssh.Context, password string) bool {
			return authByPassword(ctx.User(), password)
		}),
	)
}
