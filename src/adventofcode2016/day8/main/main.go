package main

import (
	"io/ioutil"
	"log"
	"strings"
	"adventofcode2016/day8/tasks"
	"fmt"
)

func main() {
	filename := "input.txt"
	content, err := ioutil.ReadFile(filename)
	if err != nil {
		log.Fatalf("Unable to open file %s", filename)
		return
	}
	lines := strings.Split(string(content), "\n")

	fmt.Println("Day 8 Part 1:", tasks.Day8Part1(lines, 50, 6))
	fmt.Println("Day 8 Part 2:")
	tasks.Day8Part2(lines, 50, 6)
}
