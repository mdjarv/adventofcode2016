package tasks

import (
	"regexp"
	"strconv"
)

type Disc struct {
	id        int
	positions int
	position  int
}

var re = regexp.MustCompile(`Disc #(?P<disc>[0-9]+) has (?P<positions>[0-9]+) positions; at time=0, it is at position (?P<position>[0-9]+)`)

func lineToDisc(line string) Disc {
	parts := re.FindStringSubmatch(line)
	id, _ := strconv.Atoi(parts[1])
	positions, _ := strconv.Atoi(parts[2])
	position, _ := strconv.Atoi(parts[3])
	return Disc{id, positions, position}
}

func fallsThrough(discs []Disc) bool {
	for i, disc := range discs {
		if (disc.position + 1 + i) % disc.positions != 0 {
			return false
		}
	}
	return true
}

func linesToDiscs(lines []string) []Disc {
	discs := []Disc{}

	for _, line := range lines {
		discs = append(discs, lineToDisc(line))
	}
	return discs
}

func getCorrectTime(discs []Disc) int {
	time := 0
	for fallsThrough(discs) != true {
		for i, disc := range discs {
			disc.position = (disc.position + 1) % disc.positions
			discs[i] = disc
		}
		time++
	}
	return time
}

func Day15Part1(lines []string) int {
	discs := linesToDiscs(lines)
	return getCorrectTime(discs)
}

func Day15Part2(lines []string) int {
	discs := linesToDiscs(lines)
	discs = append(discs, Disc{7, 11, 0}) // Append new disc
	return getCorrectTime(discs)
}
