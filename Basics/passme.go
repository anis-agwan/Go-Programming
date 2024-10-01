package main

import (
	"fmt"
	"os"
)

func main() {
	if len(os.Args) < 2 {
		fmt.Println("Please provide a name and a password!")
		return
	}
	if os.Args[1] == "admin" && os.Args[2] == "password" {
		fmt.Println("Welcome, admin!")
	} else {
		fmt.Println("Who are you?")
	}
}
