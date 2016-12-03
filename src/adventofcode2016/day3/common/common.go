package common

type Triangle struct {
	Side [3]int
}

func IsTrianglePossible(triangle Triangle) bool {
	if triangle.Side[0] + triangle.Side[1] <= triangle.Side[2] {
		return false
	}
	if triangle.Side[0] + triangle.Side[2] <= triangle.Side[1] {
		return false
	}
	if triangle.Side[1] + triangle.Side[2] <= triangle.Side[0] {
		return false
	}

	return true
}