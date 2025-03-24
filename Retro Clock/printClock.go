package main

import (
	"fmt"
)

func printClock(digits [10][5]string, hour int, min int, sec int, colon [5]string) {

	clock := [...][5]string {
		digits[hour/10], digits[hour%10],
		colon,
		digits[min/10], digits[min%10],
		colon,
		digits[sec/10], digits[sec%10],
	}

	alarmed := sec%10 == 0

	for line := range digits[0] {
		if alarmed {
			clock = createAlarmSign()
		}

		for i, digit := range clock {
			next := clock[i][line]
			if digit == colon && sec%2 == 0 {
				next = "   "
			}
			fmt.Print(next, " ")
		}
		fmt.Println()
	}
}
