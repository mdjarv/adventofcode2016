package tasks

import (
	"fmt"
	"regexp"
	"strings"
)

type Floors [4]Floor

type Floor struct {
	elevator                    bool
	hydrogenGenerator           bool
	hydrogenCompatibleMicrochip bool
	lithiumGenerator            bool
	lithiumCompatibleMicrochip  bool
}

var endState = Floors{
	Floor{false, false, false, false, false},
	Floor{false, false, false, false, false},
	Floor{false, false, false, false, false},
	Floor{true, true, true, true, true},
}

func printFloors(floors Floors) {
	for i := len(floors) - 1; i >= 0; i-- {
		floor := floors[i]

		e := "."
		hg := "."
		hm := "."
		lg := "."
		lm := "."

		if floor.elevator {
			e = "E"
		}

		if floor.hydrogenGenerator {
			hg = "HG"
		}

		if floor.hydrogenCompatibleMicrochip {
			hm = "HM"
		}

		if floor.lithiumGenerator {
			lg = "LG"
		}

		if floor.lithiumCompatibleMicrochip {
			lm = "LM"
		}

		fmt.Printf("F%d %-2s %-2s %-2s %-2s %-2s\n", i + 1, e, hg, hm, lg, lm)

	}
}

func addToFloor(floor Floor, content string) Floor {
	switch content {
	case "hydrogen generator":
		floor.hydrogenGenerator = true
	case "hydrogen-compatible microchip":
		floor.hydrogenCompatibleMicrochip = true
	case "lithium generator":
		floor.lithiumGenerator = true
	case "lithium-compatible microchip":
		floor.lithiumCompatibleMicrochip = true
	default:
		fmt.Println("unknown contents:", content)
	}
	return floor
}

func floorsFromLines(lines []string) Floors {
	floors := Floors{}
	for _, line := range lines {
		re := regexp.MustCompile(`The (?P<floor>\w+) floor contains (?P<contains>.*)`)
		s := re.FindStringSubmatch(line)
		currentFloor := Floor{}

		contents := strings.Split(s[2], " and ")

		for _, content := range contents {
			content = strings.TrimLeft(content, "a ")
			content = strings.TrimRight(content, ".")
			currentFloor = addToFloor(currentFloor, content)
		}

		switch (s[1]) {
		case "first":
			floors[0] = currentFloor
		case "second":
			floors[1] = currentFloor
		case "third":
			floors[2] = currentFloor
		case "fourth":
			floors[3] = currentFloor
		}
	}
	floors[0].elevator = true
	return floors
}

func findElevator(floors Floors) int {
	for i, floor := range floors {
		if floor.elevator {
			return i
		}
	}
	return -1
}

func Day11Part1(lines []string) int {
	floors := floorsFromLines(lines)
	printFloors(floors)

	moves := nextMoves(floors)


	return -1
}
func nextMoves(floors Floors) []Floors {
	result := []Floors{}

	return result
}

func Day11Part2(lines []string) int {
	return -1
}
