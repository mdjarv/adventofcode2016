package tasks

import (
	"adventofcode2016/day3/common"
)

func Part1ValidTriangleCount(lines []string) int {
	validCount := 0
	for _, line := range lines {
		if (common.IsTrianglePossible(line)) {
			validCount++
		} else {
		}
	}
	return validCount
}