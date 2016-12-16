package main

import (
	"io/ioutil"
	"log"
	"fmt"
	"strings"
	"adventofcode2016/day12/tasks"
)

func main() {
	filename := "input/day12.txt"
	content, err := ioutil.ReadFile(filename)
	if err != nil {
		log.Fatalf("Unable to open file %s", filename)
		return
	}
	lines := strings.Split(string(content), "\n")

	fmt.Println("Day 12 Part 1:", tasks.Day12Part1(lines))
	fmt.Println("Day 12 Part 2:", tasks.Day12Part2(lines))
}
