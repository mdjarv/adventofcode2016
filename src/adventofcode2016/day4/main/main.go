package main

import (
	"io/ioutil"
	"log"
	"strings"
	"adventofcode2016/day4/tasks"
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

	fmt.Println("Part 1 sum of sectors", tasks.Part1SumOfValidSectors(lines))
	fmt.Println("Part 2 northpole object storage room:", tasks.Part2GetNorthPoleObjectRoom(lines))
}
