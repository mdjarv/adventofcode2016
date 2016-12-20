package main

import (
	"fmt"
	"adventofcode2016/day15/tasks"
	"io/ioutil"
	"log"
	"strings"
)

func main() {
	filename := "input/day15.txt"
	content, err := ioutil.ReadFile(filename)
	if err != nil {
		log.Fatalf("Unable to open file %s", filename)
		return
	}
	lines := strings.Split(string(content), "\n")
	fmt.Println("Day 15 Part 1:", tasks.Day15Part1(lines))
	fmt.Println("Day 15 Part 2:", tasks.Day15Part2(lines))
}
