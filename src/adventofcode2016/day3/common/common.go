package common

import (
	"strings"
	"strconv"
)

type Triangle struct {
	Side1 int
	Side2 int
	Side3 int
}

func triangleFromString(values string) Triangle {
	sides := strings.Fields(values)
	triangle := Triangle{}
	triangle.Side1, _ = strconv.Atoi(sides[0])
	triangle.Side2, _ = strconv.Atoi(sides[1])
	triangle.Side3, _ = strconv.Atoi(sides[2])

	return triangle
}

func checkTriangleValid(triangle Triangle) bool {
	if triangle.Side1 + triangle.Side2 <= triangle.Side3 {
		return false
	}
	if triangle.Side1 + triangle.Side3 <= triangle.Side2 {
		return false
	}
	if triangle.Side2 + triangle.Side3 <= triangle.Side1 {
		return false
	}

	return true
}

func IsTrianglePossible(values string) bool {
	triangle := triangleFromString(values)
	return checkTriangleValid(triangle)
}