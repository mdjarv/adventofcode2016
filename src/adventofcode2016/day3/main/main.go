package main

import (
	"io/ioutil"
	"log"
	"fmt"
	"adventofcode2016/day3/tasks"
)

func main() {
	filename := "input.txt"
	content, err := ioutil.ReadFile(filename)
	if err != nil {
		log.Fatalf("Unable to open file %s", filename)
		return
	}

	fmt.Println("Part 1 valid triangles", tasks.Part1ValidTriangleCount(string(content)))
	fmt.Println("Part 2 valid triangles", tasks.Part2ValidTriangleCount(string(content)))
}
