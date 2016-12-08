package tasks

import (
	"fmt"
	"strings"
	"strconv"
)

func MakeRuneMatrix(width int, height int) [][]rune {
	matrix := [][]rune{}

	for y := 0; y < height; y++ {
		row := []rune{}
		for x := 0; x < width; x++ {
			row = append(row, '.')
		}
		matrix = append(matrix, row)
	}

	return matrix
}

func PrintMatrix(matrix [][]rune) {
	for _, row := range matrix {
		fmt.Println(strings.Replace(string(row), ".", " ", -1))
	}
	fmt.Println()
}

func doRect(matrix [][]rune, args string) [][]rune {
	xy := strings.Split(args, "x")
	width, _ := strconv.Atoi(xy[0])
	height, _ := strconv.Atoi(xy[1])

	for y := 0; y < height; y++ {
		for x := 0; x < width; x++ {
			matrix[y][x] = '#'
		}
	}

	return matrix
}

func copyMatrix(matrix [][]rune) [][]rune {
	newMatrix := make([][]rune, len(matrix))
	for i := range matrix {
		newMatrix[i] = make([]rune, len(matrix[i]))
		copy(newMatrix[i], matrix[i])
	}
	return newMatrix
}

func rotateColumn(matrix [][]rune, column int, amount int) [][]rune {
	buffer := copyMatrix(matrix)
	for i, row := range buffer {
		newPos := (i+amount) % len(buffer)
		matrix[newPos][column]  = row[column]
	}
	return matrix
}

func rotateRow(matrix [][]rune, row int, amount int) [][]rune {
	buffer := copyMatrix(matrix)
	for i, c := range buffer[row] {
		newPos := (i+amount) % len(buffer[row])
		matrix[row][newPos]  = c
	}
	return matrix
}

func doRotate(matrix [][]rune, args string) [][]rune {
	parts := strings.Split(args, "=")
	switch(parts[0]) {
	case "column x":
		vals := strings.Split(parts[1], " by ")
		column, _ := strconv.Atoi(vals[0])
		amount, _ := strconv.Atoi(vals[1])
		matrix = rotateColumn(matrix, column, amount)
	case "row y":
		vals := strings.Split(parts[1], " by ")
		row, _ := strconv.Atoi(vals[0])
		amount, _ := strconv.Atoi(vals[1])
		matrix = rotateRow(matrix, row, amount)
	default:
		fmt.Println("method not yet implemented", parts[0])
	}
	return matrix
}

func countLeds(matrix [][]rune) int {
	leds := 0
	for _, row := range matrix {
		for _, col := range row {
			if col == '#' {
				leds++
			}
		}
	}
	return leds
}

func Day8Part1(lines []string, width int, height int) int {
	matrix := MakeRuneMatrix(width, height)
	//fmt.Println("Start:")
	//PrintMatrix(matrix)
	for _, line := range lines {
		command := strings.SplitN(line, " ", 2)
		switch(command[0]) {
		case "rect":
			matrix = doRect(matrix, command[1])
		case "rotate":
			matrix = doRotate(matrix, command[1])
		default:
			fmt.Println("command not implemented", command[0])
			continue
		}
		//fmt.Println("Command:", line)
		//PrintMatrix(matrix)
	}
	return countLeds(matrix)
}

func Day8Part2(lines []string, width int, height int) {
	matrix := MakeRuneMatrix(width, height)
	for _, line := range lines {
		command := strings.SplitN(line, " ", 2)
		switch(command[0]) {
		case "rect":
			matrix = doRect(matrix, command[1])
		case "rotate":
			matrix = doRotate(matrix, command[1])
		default:
			fmt.Println("command not implemented", command[0])
			continue
		}
		//fmt.Println("Command:", line)
		//PrintMatrix(matrix)
	}
	PrintMatrix(matrix)
}