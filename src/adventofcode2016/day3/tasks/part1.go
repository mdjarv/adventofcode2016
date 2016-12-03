package tasks

import (
	"adventofcode2016/day3/common"
	"strings"
	"strconv"
)

func triangleFromString(values string) common.Triangle {
	sides := strings.Fields(values)
	triangle := common.Triangle{}
	triangle.Side[0], _ = strconv.Atoi(sides[0])
	triangle.Side[1], _ = strconv.Atoi(sides[1])
	triangle.Side[2], _ = strconv.Atoi(sides[2])

	return triangle
}

func Part1ValidTriangleCount(content string) int {
	lines := strings.Split(content, "\n")
	validCount := 0
	for _, line := range lines {
		triangle := triangleFromString(line)
		if (common.IsTrianglePossible(triangle)) {
			validCount++
		} else {
		}
	}
	return validCount
}