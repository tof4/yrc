package server

import (
	"log"
)

func catchFatal(err error) {
	if err != nil {
		log.Fatal(err)
	}
}
