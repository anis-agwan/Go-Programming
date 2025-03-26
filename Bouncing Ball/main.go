package main

import (
	"fmt"
	"time"
)

func main() {

	const (
		l = 10
		w = 50

		cellEmpty = " "
		cellBall  = "⚾️"
	)

	grid := make([][]bool, l) // Create outer slice of length 3
	for i := range grid {
		grid[i] = make([]bool, w) // Create inner slices of length 3
	}

	// grid[1][1] = true
	// grid[2][2] = true
	// grid[3][3] = true
	// grid[4][4] = true
	var currX = 0
	var currY = 0
	xsign := 1
	ysign := 1
	grid[currX][currY] = true
	for {
		fmt.Print("\033[H\033[2J")
		// fmt.Println("LOOK AT THAT BALL BOUNCING")
		// printTopBorder(l)

		grid[currX][currY] = true

		for x := range l {
			for y := range w {
				if grid[x][y] {
					fmt.Print(cellBall)
				} else {
					fmt.Print(cellEmpty)
				}
			}
			fmt.Println()
		}

		grid[currX][currY] = false

		if currX == 0 {
			xsign = 1
		} else if currX == l-1 {
			xsign = -1
		}

		if currY == 0 {
			ysign = 1
		} else if currY == w-1 {
			ysign = -1
		}

		currX += xsign
		currY += ysign

		time.Sleep(time.Second / 20)
	}

}
