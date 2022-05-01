package common

import (
	"log"
)

func CatchFatal(err error) {
	if err == nil {
		return
	}

	log.Fatal(err)

}
