package main

import (
	"bufio"
	"io"
	"log"
	"net"
	"strings"
)

func main() {
	con, err := net.Dial("tcp", "0.0.0.0:9999")
	if err != nil {
		log.Fatalln(err)
	}
	defer con.Close()

	serverReader := bufio.NewReader(con)

	// Waiting for the client to connect
	connmessage := "Connected..."

	switch err {
	case nil:
		if _, err = con.Write([]byte(connmessage)); err != nil {
			log.Printf("failed to request connection: %v\n", err)
		}
	case io.EOF:
		log.Println("client closed the connection")
		return
	default:
		log.Printf("client error: %v\n", err)
		return
	}

	for {
		sr, err := serverReader.ReadString('\n')

		switch err {
		case nil:
			log.Println(strings.TrimSpace(sr))
			if _, err = con.Write([]byte("Recieved: " + sr)); err != nil {
				log.Printf("failed to request connection: %v\n", err)
			}
		case io.EOF:
			log.Println("server closed the connection")
			main()
		default:
			log.Printf("server error: %v\n", err)
			main()
		}
	}
}
