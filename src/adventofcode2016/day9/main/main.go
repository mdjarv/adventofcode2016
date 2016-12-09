package main

import (
	"io/ioutil"
	"log"
	"adventofcode2016/day9/tasks"
	"fmt"
)

func main() {
	filename := "input/day9.txt"
	content, err := ioutil.ReadFile(filename)
	if err != nil {
		log.Fatalf("Unable to open file %s", filename)
		return
	}
	fmt.Println("Day 9 Part 1:", tasks.Day9Part1(string(content)))
	fmt.Println("Day 9 Part 2:", tasks.Day9Part2(string(content)))
}
