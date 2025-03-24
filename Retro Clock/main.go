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
		ssec := now.Nanosecond() / 1e8;
		// fmt.Print(ssec);
		fmt.Println()
		printClock(digits, hour, min, sec, ssec, colon)
		fmt.Println()
		time.Sleep(time.Second / 10)
	}

	
}