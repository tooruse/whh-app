package main

import (
	"bufio"
	"fmt"
	"net"
	"os"
)

const (
	connHost = "localhost"
	connPort = "8080"
	connType = "tcp"
)

func hErr(err error) {
	if err != nil {
		fmt.Println("Error:", err.Error())
		os.Exit(1)
	}
}

func main() {
	fmt.Println("Connecting to " + connType + " server " + connHost + ":" + connPort)
	conn, err := net.Dial(connType, connHost+":"+connPort)
	hErr(err)

	for {
		fmt.Println("Waiting for command")

		message, _ := bufio.NewReader(conn).ReadString('\n')

		fmt.Println("Server:", message)

		conn.Write([]byte(message))
	}
}
