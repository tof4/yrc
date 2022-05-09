package main

import (
	"flag"
	"fmt"

	"github.com/tof4/yrc/pkg/database"
)

func main() {
	rootPath := flag.String("r", "ydb", "database root path")
	channelName := flag.String("c", "", "channel name")
	flag.Parse()
	database.OpenDatabase(*rootPath)

	if *channelName == "" {
		flag.PrintDefaults()
	}

	err := database.DeleteChannel(*channelName)
	if err != nil {
		fmt.Println(err)
	}
}
