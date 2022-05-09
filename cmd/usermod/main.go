package main

import (
	"flag"
	"fmt"

	"github.com/tof4/yrc/pkg/database"
)

var rootPath *string

func main() {
	rootPath = flag.String("r", "ydb", "database root path")
	addUsername := flag.String("a", "", "add user to channel")
	delUsername := flag.String("d", "", "remove user from channel")
	channelName := flag.String("c", "", "channel name")
	flag.Parse()

	if *addUsername != "" {
		addUser(*channelName, *addUsername)
	} else if *delUsername != "" {
		delUser(*channelName, *delUsername)
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

func delUser(channelName string, username string) {
	database.OpenDatabase(*rootPath)
	err := database.RemoveFromChannel(channelName, username)

	if err != nil {
		fmt.Println(err)
	}
}
