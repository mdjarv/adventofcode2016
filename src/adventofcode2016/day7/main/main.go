package main

import (
	"io/ioutil"
	"log"
	"strings"
	"adventofcode2016/day7/tasks"
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

	fmt.Println("Part 1: found", tasks.Day7Part1(lines), "addresses with TLS support")
	fmt.Println("Part 2: found", tasks.Day7Part2(lines), "addresses with SSL support")
}
