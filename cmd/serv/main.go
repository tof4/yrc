package main

import (
	"flag"

	"github.com/tof4/yrc/pkg/server"
)

func main() {
	port := flag.Int("p", 9999, "ssh server port")
	rootPath := flag.String("d", "ydb", "database root path")
	flag.Parse()
	server.Initialize(*port, *rootPath)
}
