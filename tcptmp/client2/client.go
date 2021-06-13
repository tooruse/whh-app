package main

import (
	"fmt"
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
	defer conn.Close()
	if err != nil {
		panic(err)
	}

	//greet := "hello from client"
	//conn.Write([]byte(greet))
	//time.Sleep(time.Duration(1) * time.Second)
	/*res := make([]byte, 1024)
	_, err = conn.Read(res)
	if err != nil {
		panic(err)
	}
	fmt.Println(string(res))
	conn.Write([]byte("Recieved: " + string(res)))*/

	cli(conn)
}

func cli(conn net.Conn) {
	for {
		res := make([]byte, 1024)
		_, err := conn.Read(res)
		if err != nil {
			panic(err)
		}
		fmt.Println(string(res))

		conn.Write([]byte("Running: " + string(res)))
		//exec(string(res))
	}
}

func exec(pld string) {
	//
}
