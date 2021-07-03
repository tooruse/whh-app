package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
)

func main() {

	// read file
	data, err := ioutil.ReadFile("./example.json")
	if err != nil {
		fmt.Print(err)
	}

	// define data structure
	type DayPrice struct {
		USD float32
		EUR float32
		GBP float32
	}

	// json data
	var obj DayPrice

	// unmarshall it
	err = json.Unmarshal(data, &obj)
	if err != nil {
		fmt.Println("error:", err)
	}

	// can access using struct now
	fmt.Printf("USD : %.2f\n", obj.USD)
	fmt.Printf("EUR : %.2f\n", obj.EUR)
	fmt.Printf("GBP : %.2f\n", obj.GBP)

}
