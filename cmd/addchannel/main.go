package main

import (
	"flag"
	"fmt"

	"github.com/tof4/yrc/pkg/database"
)

func main() {
	rootPath := flag.String("r", "ydb", "database root path")
	channelName := flag.String("n", "", "channelName")
	flag.Parse()

	database.OpenDatabase(*rootPath)
	err := database.CreateChannel(*channelName)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Printf("Channel %s added \n", *channelName)
}
