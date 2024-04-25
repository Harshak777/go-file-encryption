package main

import (
	"fmt"
	"os"
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
	case "encypt":
		encyptHandle()
	case "decrypt":
		decryptHandle()
	default:
		fmt.Println("Use encrypt keyword to encypt the file, decrypt keyword to decypt the file.")
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

}

func decryptHandle() {

}

func getPassword() {

}

func validatePassword() {

}

func validateFile() {

}
