package main

import (
	"io/ioutil"
	"log"
	"fmt"
	"strings"
	"adventofcode2016/day11/tasks"
)

func main() {
	filename := "input/day11.txt"
	content, err := ioutil.ReadFile(filename)
	if err != nil {
		log.Fatalf("Unable to open file %s", filename)
		return
	}
	lines := strings.Split(string(content), "\n")

	fmt.Println("Day 11 Part 1:", tasks.Day11Part1(lines))
	fmt.Println("Day 11 Part 2:", tasks.Day11Part2(lines))
}
