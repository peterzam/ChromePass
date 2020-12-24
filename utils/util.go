package utils

import (
	"fmt"
	"io/ioutil"
	"log"
	"strings"
	"syscall"

	"golang.org/x/crypto/ssh/terminal"
)

type info struct {
	File string //file string
	Path string //filepath string
	Pass string //passphrase
}

//GetInfo - get file and passphrase
func GetInfo() info {
	var credential info
	fmt.Print("File Path(Absolute/Relative) : ")
	fmt.Scanf("%s", &credential.Path)
	byteFile, err := ioutil.ReadFile(credential.Path)
	if err != nil {
		log.Println(err)
	}
	fmt.Printf("Enter Passphrase : ")
	bytePassphrase, err := terminal.ReadPassword(int(syscall.Stdin))
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println()
	credential.File = string(byteFile)
	credential.Pass = string(bytePassphrase)
	return credential
}

//ShowCred - A prototype for show a cred
func ShowCred(data string) {
	string := strings.Split(data, ",")
	fmt.Println("\nName: " + string[0])
	fmt.Println("Url : " + string[1])
	fmt.Println("User: " + string[2])
	fmt.Println("Pass: " + string[3])
}
