package tasks

import (
	"adventofcode2016/day2/common"
)

var Part2Keypad [][]rune = [][]rune{
	[]rune{' ', ' ', '1', ' ', ' '},
	[]rune{' ', '2', '3', '4', ' '},
	[]rune{'5', '6', '7', '8', '9'},
	[]rune{' ', 'A', 'B', 'C', ' '},
	[]rune{' ', ' ', 'D', ' ', ' '},
}

var Part2InitialPosition common.Position = common.Position{X: 0, Y: 2}
