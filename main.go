package main

import (
	"bytes"
	"fmt"
	"os"

	"github.com/samrityy/file-encryption/filecrypt"
	"golang.org/x/term"
)

func main() {
	if len(os.Args) < 2 {
		printHelp()
		os.Exit(0)
	}
	function := os.Args[1]

	switch function {
	case "help":
		printHelp()
	case "encrypt":
		encryptHandle()
	case "decrypt":
		decryptHandle()
	default:
		fmt.Println("Run encrypt to encrypt a file and decrypt to decrypt  a file ")
		os.Exit(1)
	}
}

func printHelp() {
	fmt.Println("file encryption")
	fmt.Println("Simple file encryprion for your day-to-day needs")
	fmt.Println("")
	fmt.Println("Usage:")
	fmt.Println("")
	fmt.Println("\t go run . encrypt a file given a password")
	fmt.Println("\t decrypt \t tries to decrypt a file using a password")
	fmt.Println("\t help \t\t Display help text")
	fmt.Println("")
}
func encryptHandle() {
	if len(os.Args) < 3 {
		println("missing teh path to the file .For more info, run the help command")
		os.Exit(0)
	}
	file := os.Args[2]
	if !validateFile(file) {
		panic("File not found")
	}
	password := getPassword()
	fmt.Println("\nEcrypting....")
	filecrypt.Encrypt(file, password)
	fmt.Println("File encrypted successfully")

}
func decryptHandle() {
	if len(os.Args) < 3 {
		fmt.Println("missing the path to the file . for more info , run the help command")
		os.Exit(0)
	}
	file := os.Args[2]
	if !validateFile(file) {
		panic("File not found")
	}
	fmt.Println("Enter passsword")
	password, _ := term.ReadPassword(0)
	fmt.Println("\n Decrypting...")
	filecrypt.Decrypt(file, password)
	fmt.Println("\n file successfully decrypted")
}
func getPassword() []byte {
	fmt.Print("Enter password: ")
	password, _ := term.ReadPassword(0)
	if len(password) < 8 {
		fmt.Print("\n Password must be at least 8 characters long, please try again\n")
		return getPassword()
	}
	fmt.Print("/n Confirm Password")
	password2, _ := term.ReadPassword(0)
	if !validatePassword(password, password2) {
		fmt.Print("\n Password doo not match , please try again\n")
		return getPassword()
	}
	return password
}

func validatePassword(password1 []byte, password2 []byte) bool {
	return bytes.Equal(password1, password2)
}

func validateFile(file string) bool {
	if _, err := os.Stat(file); os.IsNotExist(err) {
		return false
	}
	return true
}
