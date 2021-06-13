package main

import (
	"bufio"
	"io"
	"log"
	"net"
	"os"
	"strings"
)

func main() {
	listener, err := net.Listen("tcp", "0.0.0.0:9999")
	if err != nil {
		log.Fatalln(err)
	}
	defer listener.Close()

	for {
		con, err := listener.Accept()
		if err != nil {
			log.Println(err)
			continue
		}

		// If you want, you can increment a counter here and inject to handleClientRequest below as client identifier
		go handleClientRequest(con)
	}
}

func handleClientRequest(con net.Conn) {
	defer con.Close()

	cmdReader := bufio.NewReader(os.Stdin)
	clientReader := bufio.NewReader(con)

	// Waiting for the client connection
	clientRequest, err := clientReader.ReadString('\n')

	switch err {
	case nil:
		clientRequest := strings.TrimSpace(clientRequest)
		log.Println(clientRequest)
	case io.EOF:
		log.Println("client closed the connection by terminating the process")
		return
	default:
		log.Printf("error: %v\n", err)
		return
	}

	for {
		// Responding to the client request
		pld, err := cmdReader.ReadString('\n')
		if _, err = con.Write([]byte(pld)); err != nil {
			log.Printf("failed to respond to client: %v\n", err)
		}
	}
}
