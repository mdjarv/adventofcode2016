package main

import (
	"fmt"
	"adventofcode2016/day18/tasks"
)

func main() {
	fmt.Println("Day 18 Part 1:", tasks.Day18Part1(".^.^..^......^^^^^...^^^...^...^....^^.^...^.^^^^....^...^^.^^^...^^^^.^^.^.^^..^.^^^..^^^^^^.^^^..^", 40))
	fmt.Println("Day 18 Part 2:", tasks.Day18Part1(".^.^..^......^^^^^...^^^...^...^....^^.^...^.^^^^....^...^^.^^^...^^^^.^^.^.^^..^.^^^..^^^^^^.^^^..^", 400000))
}
