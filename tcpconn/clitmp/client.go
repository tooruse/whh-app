package main

import (
	"bufio"
	"fmt"
	"net"
	"time"
)

const (
	//connection info
	connHost = "localhost"
	connPort = "8080"
	connType = "tcp"
	//command list
	exi   = "exit"
	hello = "hello"
)

func hErr(err error) {
	if err != nil {
		fmt.Println("Error:", err.Error())
	}
}

func main() {
	fmt.Println("Connecting to " + connType + " server " + connHost + ":" + connPort)
	conn, err := net.Dial(connType, connHost+":"+connPort)
	if err != nil {
		hErr(err)
		time.Sleep(10 * time.Second)
		main()
	}

	for {
		fmt.Println("Waiting for command")

		message, err := bufio.NewReader(conn).ReadString('\n')
		if err != nil {
			conn.Close()
			hErr(err)
			main()
		}

		fmt.Println(message)
		if message == hello {
			conn.Write([]byte("hello 2 you 2"))
		} else {
			conn.Write([]byte("Unknown: " + message))
		}
	}
}
