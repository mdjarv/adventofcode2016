package tasks

import (
	"fmt"
	"strconv"
)

func isWall(x int, y int, favouriteNumber int) bool {
	number := x*x + 3*x + 2*x*y + y + y*y + favouriteNumber
	bin := strconv.FormatInt(int64(number), 2)

	odd := false

	for _, r := range bin {
		if r == '1' {
			odd = !odd
		}
	}
	//fmt.Printf("input %dx%d output %b wall=%t\n", x, y, number, odd)
	return odd
}

func Day13Part1(favouriteNumber int) int {
	fmt.Println("  0123456789")
	for y := 0; y < 7; y++ {
		fmt.Printf("%d ", y)
		for x := 0; x < 10; x++ {
			if isWall(x, y, favouriteNumber) {
				fmt.Printf("#")
			} else {
				fmt.Printf(".")
			}
		}
		fmt.Println()
	}
	return -1
}

func Day12Part2(favouriteNumber int) int {
	return -1
}
