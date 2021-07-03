package main

import (
	"fmt"
	"log"
	"os/user"
)

func main() {

	user, err := user.Current()
	if err != nil {
		log.Fatalf(err.Error())
	}

	udir := user.HomeDir

	fmt.Println(udir)
	ddir := "\\AppData\\Roaming\\Microsoft\\Windows\\Start Menu\\Programs\\Startup\\lol.exe"
	fmt.Println(udir + ddir)
}
