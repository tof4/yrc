package server

import (
	"errors"

	"github.com/google/uuid"
)

type user struct {
	name         string
	passwordHash string
	clients      []client
}

func getClient(user user, clientId uuid.UUID) (client, error) {
	for _, x := range user.clients {
		if x.id == clientId {
			return x, nil
		}
	}

	return client{}, errors.New("Client not found")
}

func addClient(user *user, client client) {
	user.clients = append(user.clients, client)
}
