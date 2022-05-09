package main

import (
	"flag"

	"github.com/tof4/yrc/pkg/database"
)

func main() {
	rootPath := flag.String("r", "ydb", "database root path")
	username := flag.String("u", "", "username")
	flag.Parse()
	database.OpenDatabase(*rootPath)

	if *username == "" {
		flag.PrintDefaults()
	}

	database.DeleteUser(*username)
}
