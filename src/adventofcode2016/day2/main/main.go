package main

import (
	"io/ioutil"
	"log"
	"strings"
	"adventofcode2016/day2/common"
	"fmt"
	"adventofcode2016/day2/tasks"
)

func main() {
	filename := "input.txt"
	content, err := ioutil.ReadFile(filename)
	if err != nil {
		log.Fatalf("Unable to open file %s", filename)
		return
	}
	lines := strings.Split(string(content), "\n")

	code := common.GetCode(lines, tasks.Part1Keypad, tasks.Part1InitialPosition)
	fmt.Printf("Part 1 code is: %s\n", code)

	code = common.GetCode(lines, tasks.Part2Keypad, tasks.Part2InitialPosition)
	fmt.Printf("Part 2 code is: %s\n", code)
}

