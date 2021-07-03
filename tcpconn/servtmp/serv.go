package main

import (
	"bufio"
	"fmt"
	"net"
	"os"
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
	fmt.Println("Starting " + connType + " server on " + connHost + ":" + connPort)
	l, err := net.Listen(connType, connHost+":"+connPort)
	hErr(err)
	defer l.Close()

	for {
		c, err := l.Accept()
		hErr(err)

		fmt.Println("Client " + c.RemoteAddr().String() + " connected.")

		go handleConnection(c)
	}
}

func handleConnection(conn net.Conn) {

	reader := bufio.NewReader(os.Stdin)
	fmt.Print("PS> ")
	input, _ := reader.ReadString('\n')
	if input == exi {
		conn.Write([]byte(exi))
		os.Exit(1)
	} else if input == hello {
		conn.Write([]byte(hello))
	} else {
		conn.Write([]byte(input))
	}

	buffer, err := bufio.NewReader(conn).ReadBytes('\n')
	hErr(err)
	fmt.Print(string(buffer[:len(buffer)-1]) + "\n")

	handleConnection(conn)
}
