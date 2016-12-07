package main

import (
	"io/ioutil"
	"log"
	"strings"
	"adventofcode2016/day6/tasks"
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

	tasks.Day7Part1(lines)
}
