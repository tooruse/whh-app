package main

import (
	"encoding/base64"
	"fmt"
	"log"
	"os"
)

func hErr(err error) {
	if err != nil {
		fmt.Println(err)
	}
}

//C:\Users\lab\AppData\Roaming\Microsoft\Windows\Start Menu\Programs\Startup\
func main() {
	pld := "UXpwY1hGVnpaWEp6WEZ4c1lXSmNYRUZ3Y0VSaGRHRmNYRkp2WVcxcGJtZGNYRTFwWTNKdmMyOW1kRnhjVjJsdVpHOTNjMXhjVTNSaGNuUWdUV1Z1ZFZ4Y1VISnZaM0poYlhOY1hGTjBZWEowZFhCY1hHeHZiQT09"
	det, err := base64.StdEncoding.DecodeString(pld)
	hErr(err)
	dest, err := base64.StdEncoding.DecodeString(string(det))
	hErr(err)
	os.Mkdir(string(dest), 0755)
	if err != nil {
		log.Fatal(err)
	}
}
