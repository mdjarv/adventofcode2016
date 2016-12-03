package common

import "fmt"

type Position struct {
	X int
	Y int
}

func min(x, y int) int {
	if x < y {
		return x
	}
	return y
}

func max(x, y int) int {
	if x > y {
		return x
	}
	return y
}

func printMatrix(matrix [][]rune) {
	for _, row := range matrix {
		fmt.Print("|")
		for _, col := range row {
			fmt.Printf(" %c |", col)
		}
		fmt.Println()
	}
}

func GetCode(lines []string, keypad [][]rune, position Position) string {
	code := []rune{}

	for _, line := range lines {
		position = processLine(line, position, keypad)
		code = append(code, keypad[position.Y][position.X])
	}

	return string(code)
}

func processLine(line string, position Position, keypad [][]rune) Position {
	for i := range line {
		newPosition := position
		switch line[i] {
		case 'U':
			newPosition.Y = max(0, newPosition.Y - 1)
		case 'D':
			newPosition.Y = min(len(keypad) - 1, newPosition.Y + 1)
		case 'L':
			newPosition.X = max(0, newPosition.X - 1)
		case 'R':
			newPosition.X = min(len(keypad[0]) - 1, newPosition.X + 1)
		}
		if keypad[newPosition.Y][newPosition.X] != ' ' {
			position = newPosition
		}
	}
	return position
}