package main

import (
	"flag"
	"fmt"

	"github.com/tof4/yrc/pkg/database"
)

func main() {
	rootPath := flag.String("d", "ydb", "database root path")
	username := flag.String("u", "", "username")
	password := flag.String("p", "", "password")
	flag.Parse()

	database.OpenDatabase(*rootPath)
	err := database.CreateUser(*username, *password)

	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Printf("User %s added \n", *username)
}
