package main

import (
	"bufio"
	"errors"
	"fmt"
	"net"
)

func main() {
	fmt.Println("Starting YRC")

	PORT := ":9999"

	var activeConnections []net.Conn

	listener, err := net.Listen("tcp", PORT)
	if err != nil {
		fmt.Println(err)
		return
	}
	defer listener.Close()

	for {
		connection, err := listener.Accept()
		if err != nil {
			fmt.Println(err)
			return
		}

		go handleConnection(connection, &activeConnections)
	}
}

func handleConnection(connection net.Conn, activeConnections *[]net.Conn) {
	fmt.Printf("New connection: %s\n", connection.RemoteAddr().String())
	*activeConnections = append(*activeConnections, connection)
	fmt.Println(len(*activeConnections))

	reader := bufio.NewReader(connection)

	for {
		message, err := reader.ReadString('\n')
		if err != nil {

			if errors.As(err, &bufio.ErrFinalToken) {
				connection.Close()
				break
			} else {
				fmt.Println(err)
			}
			return
		}

		handleMessage(string(message), activeConnections)
	}
	fmt.Printf("Disconnected: %s\n", connection.RemoteAddr().String())
}

func handleMessage(message string, activeConnections *[]net.Conn) {
	for _, c := range *activeConnections {
		c.Write([]byte(message))
	}
}
