package main

import (
	"fmt"
)

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
