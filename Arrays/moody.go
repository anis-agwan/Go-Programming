// Copyright © 2018 Inanc Gumus
// Learn Go Programming Course
// License: https://creativecommons.org/licenses/by-nc-sa/4.0/
//
// For more tutorials  : https://learngoprogramming.com
// In-person training  : https://www.linkedin.com/in/inancgumus/
// Follow me on twitter: https://twitter.com/inancgumus

package main

import (
	"math/rand"
	"os"
	"fmt"
	"time"
	"os/exec"
	"runtime"
)

// ---------------------------------------------------------
// EXERCISE: Moodly
//
//   1. Get username from command-line
//
//   2. Display the usage if the username is missing
//
//   3. Create an array
//     1. Add three positive mood messages
//     2. Add three negative mood messages
//
//   4. Randomly select and print one of the mood messages
//
// EXPECTED OUTPUT
//
//   go run main.go
//     [your name]
//
//   go run main.go Socrates
//     Socrates feels good 👍
//
//   go run main.go Socrates
//     Socrates feels bad 👎
//
//   go run main.go Socrates
//     Socrates feels sad 😞
//
//   go run main.go Socrates
//     Socrates feels happy 😀
//
//   go run main.go Socrates
//     Socrates feels awesome 😎
//
//   go run main.go Socrates
//     Socrates feels terrible 😩
// ---------------------------------------------------------

// clearScreen clears the terminal screen
func clearScreen() {
	var cmd *exec.Cmd
	if runtime.GOOS == "windows" {
		cmd = exec.Command("cmd", "/c", "cls")
	} else {
		cmd = exec.Command("clear") // for linux and macOS
	}
	cmd.Stdout = os.Stdout
	cmd.Run()
}

func main() {
	// Clear the screen before printing
	clearScreen()
	
	now := time.Now()
	hour := now.Hour()
	
	// Demonstrate types
	fmt.Printf("Type of now: %T\n", now)           // time.Time
	fmt.Printf("Type of hour: %T\n", hour)         // int
	fmt.Printf("Value of hour: %d\n", hour)        // e.g., 14
	
	// Other time components and their types
	fmt.Printf("Type of minute: %T\n", now.Minute())   // int
	fmt.Printf("Type of second: %T\n", now.Second())   // int
	fmt.Printf("Type of nanosecond: %T\n", now.Nanosecond()) // int

	if(len(os.Args) < 2) {
		fmt.Println("anis")
		return
	}

	var username string = os.Args[1]
	var moods = [6]string {
		"feels happy 😀",
		"feels good 👍",
		"feels awesome 😎",
		"feels sad 😞",
		"feels bad 👎",
		"feels terrible 😩",
	}

	rand.Seed(time.Now().UnixNano())
	n := rand.Intn(len(moods))

	fmt.Printf("%s feels %s\n", username, moods[n])

}