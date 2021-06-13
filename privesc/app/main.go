package main

import (
	"fmt"
	"time"
)

func main() {
	fmt.Println("Hello world")
	time.Sleep(5 * time.Second)
	fmt.Println("Goodbye")
	time.Sleep(1 * time.Second)
}
