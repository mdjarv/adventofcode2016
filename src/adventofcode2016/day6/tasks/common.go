package tasks

import "bytes"

func StringsToRuneMatrix(lines []string) [][]rune {
	matrix := [][]rune{}
	for i, line := range lines {
		matrix = append(matrix, []rune{})
		for _, c := range line {
			matrix[i] = append(matrix[i], c)
		}
	}
	return matrix
}

func ColumnsToStrings(lines []string) []string {
	matrix := StringsToRuneMatrix(lines)

	var buffers = []bytes.Buffer{}
	for i, _ := range matrix[0] {
		buffers = append(buffers, bytes.Buffer{})
		for j, _ := range matrix {
			buffers[i].WriteRune(matrix[j][i])
		}
	}

	result := []string{}
	for _, row := range buffers {
		result = append(result, row.String())
	}

	return result
}

func CountCharacters(word string) map[rune]int {
	var counters = make(map[rune]int)
	for _, r := range word {
		_, ok := counters[r]
		if ok {
			counters[r]++
		} else {
			counters[r] = 1
		}
	}
	return counters
}