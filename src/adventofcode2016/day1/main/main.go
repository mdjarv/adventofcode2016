package main

import (
	"io/ioutil"
	"log"
	"adventofcode2016/day1/tasks"
	"fmt"
)

func main() {
	filename := "input.txt"
	content, err := ioutil.ReadFile(filename)
	if err != nil {
		log.Fatalf("Unable to open file %s", filename)
		return
	}
	fmt.Println("Part 1 distance", tasks.GetPart1Distance(string(content)))
	fmt.Println("Part 2 distance", tasks.GetPart2Distance(string(content)))
}
