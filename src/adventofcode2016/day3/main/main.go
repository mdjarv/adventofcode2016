package main

import (
	"io/ioutil"
	"log"
	"fmt"
	"strings"
	"adventofcode2016/day3/tasks"
)

func main() {
	filename := "input.txt"
	content, err := ioutil.ReadFile(filename)
	if err != nil {
		log.Fatalf("Unable to open file %s", filename)
		return
	}
	lines := strings.Split(string(content), "\n")

	fmt.Println("Part 1 valid triangles", tasks.Part1ValidTriangleCount(lines))
}
