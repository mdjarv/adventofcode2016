package main

import (
	"fmt"
	"adventofcode2016/day19/tasks"
	"flag"
	"os"
	"log"
	"runtime/pprof"
)

var cpuprofile = flag.String("cpuprofile", "", "write cpu profile to file")

func main() {
	flag.Parse()
	if *cpuprofile != "" {
		f, err := os.Create(*cpuprofile)
		if err != nil {
			log.Fatal(err)
		}
		pprof.StartCPUProfile(f)
		defer pprof.StopCPUProfile()
	}
	fmt.Println("Day 19 Part 1:", tasks.Day19Part1(3005290))
	fmt.Println("Day 19 Part 2:", tasks.Day19Part2(3005290))
}
