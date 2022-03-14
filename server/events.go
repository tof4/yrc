package main

import "fmt"

func eventJoin(client *yrcClient) {
	broadcast(*client, fmt.Sprintf("join/%d/%s", client.id, client.nickname))
}

func eventMessage(client *yrcClient, message string) {
	broadcast(*client, fmt.Sprintf("message/%d/%s", client.id, message))
}
