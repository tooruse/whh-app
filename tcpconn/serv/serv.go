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
	fmt.Print("CMD> ")
	input, _ := reader.ReadString('\n')
	conn.Write([]byte(input))

	buffer, err := bufio.NewReader(conn).ReadBytes('\n')
	hErr(err)
	fmt.Println("Client:", string(buffer[:len(buffer)-1]))

	handleConnection(conn)
}
