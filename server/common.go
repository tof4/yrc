package server

import "log"

func catch(err error) {
	if err != nil {
		log.Fatal(err)
	}
}
