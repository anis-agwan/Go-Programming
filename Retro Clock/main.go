package main

import (
	"fmt"
	// "os"
	"time"
)

func createPlaceholder() [10][5]string {
	type placeholder [5]string

	zero := placeholder {
		"███",
		"█ █",
		"█ █",
		"█ █",
		"███",
	}

	one := placeholder {
		"██ ",
		" █ ",
		" █ ",
		" █ ",
		"███",
	}

	two := placeholder {
		"███",
		"  █",
		"███",
		"█  ",
		"███",
	}

	three := placeholder {
		"███",
		"  █",
		"███",
		"  █",
		"███",
	}

	four := placeholder {
		"█ █",
		"█ █",
		"███",
		"  █",
		"  █",
	}
	five := placeholder {
		"███",
		"█  ",
		"███",
		"  █",
		"███",
	}
	six := placeholder {
		"███",
		"█  ",
		"███",
		"█ █",
		"███",
	}
	seven := placeholder {
		"███",
		"  █",
		"  █",
		"  █",
		"  █",
	}
	eight := placeholder {
		"███",
		"█ █",
		"███",
		"█ █",
		"███",
	}
	nine := placeholder {
		"███",
		"█ █",
		"███",
		"  █",
		"███",
	}

	digits := [10][5]string{
		zero, one, two, three, four, five, six, seven, eight, nine,
	}

	return digits
}

func printClock(digits [10][5]string, hour int, min int, sec int, colon [5]string) {
	type placeholder [5]string

	clock := [...]placeholder {
		digits[hour/10], digits[hour%10],
		colon,
		digits[min/10], digits[min%10],
		colon,
		digits[sec/10], digits[sec%10],
	}

	for line := range digits[0] {
		for digit := range clock {
			fmt.Print(clock[digit][line], " ")
		}
		fmt.Println()
	}
}

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