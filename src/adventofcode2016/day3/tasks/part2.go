package tasks

import (
	"adventofcode2016/day3/common"
	"strings"
	"strconv"
)

func matrixToTriangles(matrix [3][3]int) [3]common.Triangle {
	triangles := [3]common.Triangle{}

	for i := range matrix {
		for j := range matrix[i] {
			triangles[j].Side[i] = matrix[i][j]
		}
	}
	return triangles
}

func Part2ValidTriangleCount(content string) int {
	lines := strings.Split(content, "\n")
	triangles := []common.Triangle{}

	block := [3][3]int{}

	for i := range lines {
		if (i > 0 && i % 3 == 0) {
			for _, t := range matrixToTriangles(block) {
				triangles = append(triangles, t)
			}
		}
		s := strings.Fields(lines[i])
		block[i % 3][0], _ = strconv.Atoi(s[0])
		block[i % 3][1], _ = strconv.Atoi(s[1])
		block[i % 3][2], _ = strconv.Atoi(s[2])
	}
	for _, t := range matrixToTriangles(block) {
		triangles = append(triangles, t)
	}

	validCount := 0
	for _, triangle := range triangles {
		if (common.IsTrianglePossible(triangle)) {
			validCount++
		}
	}
	return validCount
}