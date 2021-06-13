package main

import (
	"bufio"
	"fmt"
	"log"
	"net"
	"os"
)

const (
	host  = "localhost"
	port  = "3333"
	prtcl = "tcp"
)

func main() {
	// Listen for incoming connections.
	l, err := net.Listen(prtcl, host+":"+port)
	if err != nil {
		fmt.Println("Error listening:", err.Error())
		os.Exit(1)
	}
	// Close the listener when the application closes.
	defer l.Close()

	fmt.Println("Listening on " + host + ":" + port)
	for {
		// Listen for an incoming connection.
		conn, err := l.Accept()
		if err != nil {
			fmt.Println("Error accepting: ", err.Error())
			os.Exit(1)
		}
		// Handle connections in a new goroutine.
		go cmdclient(conn)
	}
}

// Handles incoming requests.
func cmdclient(conn net.Conn) {
	//greet(conn)
	cmd(conn)
	conn.Close()
}

func cmd(conn net.Conn) {
	fmt.Println("New Connection...")
	cmdReader := bufio.NewReader(os.Stdin)

	for {
		//Response
		res := make([]byte, 1024)
		_, err := conn.Read(res)
		if err != nil {
			panic(err)
		}
		fmt.Println(string(res))

		//CMD to send
		pld, err := cmdReader.ReadString('\n')
		if _, err = conn.Write([]byte(pld)); err != nil {
			log.Printf("failed to send to client: %v\n", err)
		}

	}
}
