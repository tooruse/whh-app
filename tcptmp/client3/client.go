package main

import (
	"fmt"
	"log"
	"net"
)

// TCPClient struct
type TCPClient struct {
	Host string
	Port int
}

const (
	servip = "localhost"
	port   = "3333"
	prtcl  = "tcp"
)

// Start TCPClient
func main() {
	conn, err := net.Dial(prtcl, servip+":"+port)
	handleErr(err)
	defer conn.Close()

	// write to socket
	message := "hello world"
	conn.Write([]byte(message))

	// read from socket
	// (assume response is 'Test\n')
	reply := make([]byte, 1024)
	conn.Read(reply)
	fmt.Println(string(reply))

	resp := "Return: " + string(reply)
	conn.Write([]byte(resp))

	reply = make([]byte, 1024)
	conn.Read(reply)
	fmt.Println(string(reply))

	resp = "Return: " + string(reply)
	conn.Write([]byte(resp))
}

func exec(pld string) {
	//
}

func handleErr(err error) {
	if err != nil {
		log.Println(err)
	}
}
