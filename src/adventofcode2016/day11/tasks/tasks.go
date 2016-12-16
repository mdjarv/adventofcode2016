package tasks

import (
	"fmt"
	"regexp"
	"strings"
)

type State struct {
	elevator int
	floor    [4]byte
	previousState *State
}

var endState = State{
	elevator: 3,
	floor: [4]byte{0x00, 0x00, 0x00, 0x0f},
}

func printState(state State) {
	for i := len(state.floor) - 1; i >= 0; i-- {
		printFloor(state, i)
	}
	fmt.Println()
}

func addToFloor(state State, floor int, content string) State {
	switch content {
	case "hydrogen generator":
		state.floor[floor] |= 0x08
	case "hydrogen-compatible microchip":
		state.floor[floor] |= 0x04
	case "lithium generator":
		state.floor[floor] |= 0x02
	case "lithium-compatible microchip":
		state.floor[floor] |= 0x01
	}
	return state
}

func stateFromLines(lines []string) State {
	state := State{}
	for _, line := range lines {
		re := regexp.MustCompile(`The (?P<floor>\w+) floor contains (?P<contains>.*)`)
		parsedLine := re.FindStringSubmatch(line)
		floor := 0
		switch (parsedLine[1]) {
		case "first":
			floor = 0
		case "second":
			floor = 1
		case "third":
			floor = 2
		case "fourth":
			floor = 3
		}

		contents := strings.Split(parsedLine[2], " and ")

		for _, content := range contents {
			content = strings.TrimLeft(content, "a ")
			content = strings.TrimRight(content, ".")
			state = addToFloor(state, floor, content)
		}

	}
	return state
}

func Day11Part1(lines []string) int {
	state := stateFromLines(lines)
	printState(state)
	steps := []State{}
	steps, found := walkState(state, steps)
	fmt.Println(found, len(steps))
	return len(steps)
}

func sameState(state1 State, state2 State) bool {
	return state1.elevator == state2.elevator && state1.floor == state2.floor
}

func walkState(state State, previousSteps []State) ([]State, bool) {
	previousSteps = append(previousSteps, state)
	fmt.Println(len(previousSteps), "walking state", state)
	if sameState(state, endState) {
		// Found solution, return
		return previousSteps, true
	} else {
		movesAvailable := nextMoves(state, previousSteps)

		for _, move := range movesAvailable {
			if sameState(move, endState) {
				previousSteps = append(previousSteps, move)
				return previousSteps, true
			}
		}


		//for _, nextState := range nextMoves(state, previousSteps) {
		//	// Possible moves
		//	nextPrevSteps := append(previousSteps, state)
		//	nextPrevSteps, found := walkState(nextState, nextPrevSteps)
		//	if found {
		//		return nextPrevSteps, found
		//	}
		//}
	}
	return previousSteps, false
}
func unvisitedState(state State, visitedStates []State) bool {
	for _, old := range visitedStates {
		if sameState(state, old) {
			return false
		}
	}
	return true
}

func nextMoves(state State, previousStates []State) []State {
	result := []State{}

	// Possible moves with elevator
	possibleFloors := []int{}
	if state.floor[state.elevator] == 0x00 {
		// Empty floor, elevator will not be able to move
		return result
	} else if state.elevator == 0 {
		possibleFloors = append(possibleFloors, 1)
	} else if state.elevator == 3 {
		possibleFloors = append(possibleFloors, 2)
	} else {
		possibleFloors = append(possibleFloors, state.elevator - 1)
		possibleFloors = append(possibleFloors, state.elevator + 1)
	}

	for _, destinationFloor := range possibleFloors {
		for _, v := range getPossibleInventory(state) {
			//fmt.Printf("  possible inventory %04b\n", v)
			newState := State{elevator:destinationFloor, floor:state.floor}

			newState.floor[state.elevator] = newState.floor[state.elevator] - v
			newState.floor[newState.elevator] = newState.floor[newState.elevator] + v

			if validState(newState) && unvisitedState(newState, previousStates) {
				result = append(result, newState)
			}
		}
	}

	return result
}
func validState(state State) bool {
	for _, floor := range state.floor {
		for _, f := range invalidFloor {
			if floor == f {
				return false
			}
		}
	}
	return true
}

func getPossibleInventory(state State) []byte {
	possibleInventory := []byte{}
	for _, inv := range possibleElevatorContents {
		if inv & state.floor[state.elevator] == inv {
			possibleInventory = append(possibleInventory, inv)
		}
	}
	return possibleInventory
}

var invalidFloor = []byte{
	0x09, // 1001
	0x06, // 0110
}

var possibleElevatorContents = []byte{
	0x01, // 0001
	0x03, // 0011
	0x05, // 0101
	0x09, // 1001
	0x02, // 0010
	0x06, // 0110
	0x0A, // 1010
	0x04, // 0100
	0x0C, // 1100
	0x08, // 1000
}

func printFloor(state State, floor int) {
	e := "."
	hg := "."
	hm := "."
	lg := "."
	lm := "."

	if state.elevator == floor {
		e = "E"
	}

	if state.floor[floor] & 0x08 != 0 {
		hg = "HG"
	}

	if state.floor[floor] & 0x04 != 0 {
		hm = "HM"
	}

	if state.floor[floor] & 0x02 != 0 {
		lg = "LG"
	}

	if state.floor[floor] & 0x01 != 0 {
		lm = "LM"
	}

	fmt.Printf("F%d %-2s %-2s %-2s %-2s %-2s\n", floor + 1, e, hg, hm, lg, lm)
}

func Day11Part2(lines []string) int {
	return len(lines)
}
