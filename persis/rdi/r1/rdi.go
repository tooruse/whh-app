package main

import (
	"fmt"
	"io/ioutil"
	"os"
)

func hErr(err error) {
	if err != nil {
		fmt.Println(err)
	}
}

func main() {
	fmt.Println("Hello World")
}

func mkd() {
	newFile, err := os.Create("lol.dll")
	hErr(err)
	newFile.Close()

	data := "arrray of strings?"

	ioutil.WriteFile("lol.dll", []byte(data), 0755)
	hErr(err)

	fmt.Println(src)
	err = os.Remove(src)
	hErr(err)
}

func ind() {

}
