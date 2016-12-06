package tasks

import (
	"bytes"
)

func Part1SignalsAndNoise(lines []string) string {
	output := bytes.Buffer{}

	for _, word := range ColumnsToStrings(lines) {
		counters := CountCharacters(word)
		currentChar := ' '
		currentCount := 0
		for character, count := range counters {
			if count > currentCount {
				currentChar = character
				currentCount = count
			}
		}
		output.WriteRune(currentChar)

	}
	return string(output.String())
}

func Part2SignalsAndNoise(lines []string) string {
	output := bytes.Buffer{}

	for _, word := range ColumnsToStrings(lines) {
		counters := CountCharacters(word)
		currentChar := ' '
		currentCount := -1
		for character, count := range counters {
			if currentCount == -1 || count < currentCount {
				currentChar = character
				currentCount = count
			}
		}
		output.WriteRune(currentChar)

	}
	return string(output.String())
}