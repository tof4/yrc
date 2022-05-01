package server

import (
	"bufio"
	"errors"
	"fmt"
	"log"

	"github.com/gliderlabs/ssh"
	"github.com/google/uuid"
	"github.com/tof4/yrc/pkg/database"
	"golang.org/x/term"
)

var clients []client

func Initialize(port int, rootPath string) {
	database.OpenDatabase(rootPath)
	log.Println(fmt.Sprintf("Starting ssh server on port %d", port))
	log.Println(fmt.Sprintf("Database set in %s", rootPath))
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
			ssh.HostKeyFile(database.Paths.Key),
		))
}

func handleConnect(session ssh.Session) {
	user, err := database.GetUser(session.User())
	if err != nil {
		return
	}

	client := client{
		session:  session,
		id:       uuid.New(),
		username: user.Name}

	clients = append(clients, client)

	log.Println(fmt.Sprintf("Connected %s(%s) from %s", client.username, client.id, client.getAddress()))

	terminal := term.NewTerminal(session, "")
	for {
		line, err := terminal.ReadLine()
		if err != nil {

			if errors.As(err, &bufio.ErrFinalToken) {
				log.Println(fmt.Sprintf("Disconnected %s(%s)", client.username, client.id))

				removeClient(client)

				break
			} else {
				log.Println(err)
			}
			return
		}

		handleInput(line, client)
	}
}

func replyWithError(client client, err error) {
	message := fmt.Sprintf("Error: %s \n", err)
	client.sendData(message)
}

func removeClient(client client) {
	for i, x := range clients {
		if x.id == client.id {
			clients[i] = clients[len(clients)-1]
			clients = clients[:len(clients)-1]
			return
		}
	}
}
