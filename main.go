package main

import (
	"bytes"
	"fmt"
	"harshak777/go-file-encrption/filecrypt"
	"os"

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
		encyptHandle()
	case "decrypt":
		decryptHandle()
	default:
		fmt.Println("Use encrypt keyword to encrypt the file, decrypt keyword to decypt the file.")
		os.Exit(1)
	}
}

func printHelp() {
	fmt.Printf("Usage: \tgo run . encrypt /path/to/your/file")
	fmt.Println("")
	fmt.Println("Commands:")
	fmt.Println("")
	fmt.Println("\t encrypt\tEncrypts a file given a password")
	fmt.Println("\t decrypt\tTries to decrypt a file using a password")
	fmt.Println("\t help\t\tDisplays help text")
	fmt.Println("")
}

func encyptHandle() {
	if len(os.Args) < 3 {
		fmt.Println("Missing the path to the file. For more info use the help command.")
		os.Exit(0)
	}

	file := os.Args[2]

	if !validateFile(file) {
		panic("File not found")
	}

	password := getPassword()

	fmt.Println("___ENCRYPTING___")

	filecrypt.Encrypt(file, password)

	fmt.Println("Encypted the file.")
}

func decryptHandle() {
	if len(os.Args) < 3 {
		fmt.Println("Missing the path to the file. For more info use the help command.")
		os.Exit(0)
	}

	file := os.Args[2]

	if !validateFile(file) {
		panic("File not found")
	}

	password := getPassword()

	fmt.Println("___DECRYPTING___")

	filecrypt.Decrypt(file, password)

	fmt.Println("Decypted the file.")
}

func getPassword() []byte {
	fmt.Printf("Enter the password. \n")
	password, _ := term.ReadPassword(0)

	fmt.Printf("Confirm the password. \n")
	confirmPassword, _ := term.ReadPassword(0)

	if !validatePassword(password, confirmPassword) {
		fmt.Println("Passwords didn't match")
		getPassword()
	}

	return password
}

func validatePassword(password []byte, confirmPassword []byte) bool {
	return bytes.Equal(password, confirmPassword)
}

func validateFile(file string) bool {
	if _, err := os.Stat(file); os.IsNotExist(err) {
		return false
	}
	return true
}
