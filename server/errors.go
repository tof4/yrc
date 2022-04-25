package server

import (
	"fmt"
	"log"
)

func catchFatal(err error) {
	if err == nil {
		return
	}

	log.Fatal(err)

}

func replyWithError(client client, err error) {
	message := fmt.Sprintf("Error: %s \n", err)
	client.sendData(message)
}
