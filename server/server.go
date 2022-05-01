package server

import (
	"bufio"
	"errors"
	"fmt"
	"log"

	"github.com/gliderlabs/ssh"
	"github.com/google/uuid"
	"golang.org/x/term"
)

func Initialize(port int, rootPath string) {
	log.Println(fmt.Sprintf("Starting ssh server on port %d", port))
	log.Println(fmt.Sprintf("Database set in %s", rootPath))
	openDatabase(rootPath)
	listenSsh(port)
}

func listenSsh(port int) {

	ssh.Handle(func(s ssh.Session) {
		handleConnect(s)
	})

	log.Fatal(
		ssh.ListenAndServe(fmt.Sprintf(":%d", port), nil,
			ssh.PasswordAuth(func(ctx ssh.Context, password string) bool {
				return authByPassword(ctx.User(), password)
			}),
			ssh.HostKeyFile(paths.key),
		))
}

func handleConnect(session ssh.Session) {
	user, err := getUser(session.User())
	if err != nil {
		return
	}

	client := client{
		session: session,
		id:      uuid.New(),
		user:    user}

	addClient(user, client)
	log.Println(fmt.Sprintf("Connected %s(%s) from %s", user.name, client.id, client.getAddress()))

	terminal := term.NewTerminal(session, "")
	for {
		line, err := terminal.ReadLine()
		if err != nil {

			if errors.As(err, &bufio.ErrFinalToken) {
				log.Println(fmt.Sprintf("Disconnected %s(%s)", user.name, client.id))
				removeClient(user, client)
				break
			} else {
				log.Println(err)
			}
			return
		}

		handleInput(line, client)
	}
}
