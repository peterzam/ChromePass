package main

import (
	"ChromePass/utils"
	"fmt"
	"log"
	"os"
	"os/exec"
	"runtime"
	"strings"
	"time"
)

var clear map[string]func()

func init() {
	clear = make(map[string]func())
	clear["linux"] = func() {
		cmd := exec.Command("clear")
		cmd.Stdout = os.Stdout
		cmd.Run()
	}
	clear["windows"] = func() {
		cmd := exec.Command("cmd", "/c", "cls")
		cmd.Stdout = os.Stdout
		cmd.Run()
	}
}

func main() {

	Options()
}

//Options - Option Menu
func Options() {
StartOption:
	var option int
	fmt.Print("\nOptions: \n" +
		"1. Encrypt File,\n" +
		"2. Decrypt File,\n" +
		"3. Get Credential from Encrypted file\n" +
		"4. Show all Credentials from Encrypted file\n" +
		"5. Add a Credential to Encrypted file\n" +
		"6. Remove a Credential from Encrypted file\n" +
		"0. Exit\n:")
	fmt.Scanf("%d", &option)

	switch option {

	case 0:
		os.Exit(0)
	case 1:
		EncryptFile()
		fmt.Println("Done!")
	case 2:
		DecryptFile()
		fmt.Println("Done!")
	case 3:
		GetCred()
		fmt.Println("Done!")
	case 4:
		GetAllCred()
		fmt.Println("Done!")
	case 5:
		AddCred()
		fmt.Println("Done!")
	case 6:
		DelCred()
		fmt.Println("Done!")
	default:
		goto StartOption
	}

}

//EncryptFile - Read file, encrypt and save
func EncryptFile() {
	credential := utils.GetInfo()
	enc := utils.AESEncrypt(credential.File, utils.StringToSHA256(credential.Pass))
	fmt.Print("Output file name : ")
	fmt.Scanf("%s", &credential.Path)
	outputFile, err := os.Create(credential.Path)
	if err != nil {
		log.Fatal(err)
	}
	outputFile.WriteString(enc)
	defer outputFile.Close()
	ClearScreen(5)
}

//DecryptFile - Read file, Decrypt and save
func DecryptFile() {
	credential := utils.GetInfo()
	dec := utils.AESDecrypt(credential.File, utils.StringToSHA256(credential.Pass))
	fmt.Print("Output file name : ")
	fmt.Scanf("%s", &credential.Path)
	outputFile, err := os.Create(credential.Path)
	if err != nil {
		log.Fatal(err)
	}
	outputFile.WriteString(dec)
	defer outputFile.Close()
	ClearScreen(5)
}

//GetCred - Search and show cred from encrypted file
func GetCred() {
	credential := utils.GetInfo()
	var found bool = false
	dec := utils.AESDecrypt(credential.File, utils.StringToSHA256(credential.Pass))
	var searchValue string
	fmt.Print("Search : ")
	fmt.Scanf("%s", &searchValue)
	for _, s := range strings.Fields(dec) {
		words := strings.FieldsFunc(s, func(r rune) bool { return strings.ContainsRune(" .,@/", r) })
		for _, w := range words {
			if w == searchValue {
				utils.ShowCred(s)
				found = true
				break
			}
		}
	}
	if found == false {
		fmt.Println("Not found!")
	}
	ClearScreen(30)
}

//GetAllCred - Search and show all creds from encrypted file
func GetAllCred() {
	const MAX int = 1024
	credential := utils.GetInfo()
	dec := utils.AESDecrypt(credential.File, utils.StringToSHA256(credential.Pass))
	for i, s := range strings.Fields(dec) {
		fmt.Printf("\n---------------------------\nNo : %d", i)
		utils.ShowCred(s)
	}
	ClearScreen(30)
}

//DelCred - delete cred from encrypted file
func DelCred() {
	credential := utils.GetInfo()
	dec := utils.AESDecrypt(credential.File, utils.StringToSHA256(credential.Pass))
	var selectNo int
	fmt.Print("No. to delete: ")
	fmt.Scanf("%d", &selectNo)
	credential.File = ""
	for i, s := range strings.Fields(dec) {
		if i != selectNo {
			credential.File = credential.File + s + "\n"
		} else {
			fmt.Println("Deleted Cred : ")
			fmt.Print("--------------------------")
			utils.ShowCred(s)
		}
	}
	enc := utils.AESEncrypt(credential.File, utils.StringToSHA256(credential.Pass))
	outputFile, err := os.Create(credential.Path)
	if err != nil {
		log.Fatal(err)
	}
	outputFile.WriteString(enc)
	defer outputFile.Close()
	ClearScreen(10)
}

//AddCred - Add a cred to encrypted file
func AddCred() {
	credential := utils.GetInfo()
	dec := utils.AESDecrypt(credential.File, utils.StringToSHA256(credential.Pass))
	var newCred string
	fmt.Print("New Cred : ")
	fmt.Scanf("%s", &newCred)
	fmt.Println("Added Cred : ")
	fmt.Print("--------------------------")
	utils.ShowCred(newCred)
	dec = dec + "\n" + newCred
	enc := utils.AESEncrypt(dec, utils.StringToSHA256(credential.Pass))
	outputFile, err := os.Create(credential.Path)
	if err != nil {
		log.Fatal(err)
	}
	outputFile.WriteString(enc)
	defer outputFile.Close()
	ClearScreen(10)
}

//ClearScreen - Clear screen for safty of decrypted creds.
func ClearScreen(t int) {
	ticker := time.Tick(time.Second)
	for i := t; i >= 0; i-- {
		<-ticker
		fmt.Printf("\rTerminating in %ds ", i)
	}
	value, ok := clear[runtime.GOOS]
	if ok {
		value()
	} else {
		panic("Your platform is unsupported! I can't clear terminal screen :(")
	}
}
