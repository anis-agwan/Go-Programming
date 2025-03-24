package main

import (
	"fmt"
	// "os"
	"time"
)


func main() {
	
	// fmt.Println(now.Minute())
	// fmt.Println(now.Second())

	digits := createPlaceholder()
	colon := [5]string {
		"   ",
		" ░ ",
		"   ",
		" ░ ",
		"   ",
	}	

	for {
		
		fmt.Print("\033[H\033[2J")  // Clear screen using ANSI escape codes
		now := time.Now()
		hour := now.Hour()
		min := now.Minute()
		sec := now.Second()
		fmt.Println()
		printClock(digits, hour, min, sec, colon)
		fmt.Println()
		time.Sleep(time.Second)
	}

	
}