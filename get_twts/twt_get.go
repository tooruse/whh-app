package main

import (
	"encoding/base64"
	"fmt"

	twt_get "github.com/n0madic/twitter-scraper"
)

func hErr(err error) {
	if err != nil {
		fmt.Println(err)
	}
}

func main() {

	id := "1403851092429991941"
	scraper := twt_get.New()
	tweet, err := scraper.GetTweet(id)
	hErr(err)
	data, err := base64.StdEncoding.DecodeString(tweet.Text)
	hErr(err)
	fmt.Println(string(data))
}
