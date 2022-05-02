package main

import (
	"flag"
	"fmt"

	"github.com/tof4/yrc/pkg/database"
)

var rootPath *string

func main() {
	rootPath = flag.String("r", "ydb", "database root path")
	channelName := flag.String("c", "", "channel name")
	addUsername := flag.String("a", "", "add user")
	delUsername := flag.String("d", "", "delete user")
	flag.Parse()

	if *addUsername != "" {
		addUser(*channelName, *addUsername)
	} else if *delUsername != "" {
		fmt.Println("Not implemented")
	} else {
		flag.PrintDefaults()
	}
}

func addUser(channelName string, username string) {
	database.OpenDatabase(*rootPath)
	err := database.AddToChannel(channelName, username)

	if err != nil {
		fmt.Println(err)
	}
}
