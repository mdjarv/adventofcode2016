package main

import (
	"io/ioutil"
	"log"
	"adventofcode2016/day10/tasks"
	"fmt"
	"strings"
)

func main() {
	filename := "input/day10.txt"
	content, err := ioutil.ReadFile(filename)
	if err != nil {
		log.Fatalf("Unable to open file %s", filename)
		return
	}
	lines := strings.Split(string(content), "\n")
	fmt.Println("Day 10 Part 1:", tasks.Day10Part1(lines, 61, 17))
	fmt.Println("Day 10 Part 2:", tasks.Day10Part2(lines, 61, 17))
}
