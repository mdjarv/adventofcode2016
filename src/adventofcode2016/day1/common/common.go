package common

import (
	"strings"
	"strconv"
)

const (
	North = iota
	East = iota
	South = iota
	West = iota
)

type Position struct {
	X int
	Y int
	Direction int
}

func abs(value int) int {
	mask := value >> 31
	value = value ^ mask
	return value - mask
}

func Distance(p1 Position, p2 Position) int {
	x := p2.X - p1.X
	y := p2.Y - p1.Y
	return abs(x) + abs(y)
}

func GetPath(commands string) []Position {
	positions := []Position{{0,0, North}}
	for _, command := range strings.Split(commands, ", ") {
		positions = doStep(command, positions)
	}
	return positions
}

func doStep(cmd string, positions []Position) []Position {
	position := positions[len(positions)-1]
	switch cmd[0] {
	case 'L':
		position.Direction = position.Direction - 1
	case 'R':
		position.Direction = position.Direction + 1
	}
	if position.Direction < North {
		position.Direction = West
	} else if position.Direction > West {
		position.Direction = North
	}

	steps, _ := strconv.Atoi(cmd[1:])
	for i := 0; i < steps; i++ {
		switch(position.Direction) {
		case North:
			position.Y++
		case East:
			position.X++
		case South:
			position.Y--
		case West:
			position.X--
		}
		positions = append(positions, position)
	}

	return positions
}