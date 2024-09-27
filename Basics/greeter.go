package main

import (
	"fmt"
	"os"
)

func main() {
	
	var Args []string
	Args = os.Args

	fmt.Println("Hello, " + Args[1] + "!")
}