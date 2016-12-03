package common

import "testing"

type testdata struct {
	position1 Position
	position2 Position
	distance int
}

var tests = []testdata{
	{Position{0, 0}, Position{0, 2}, 2},
	{Position{0, 0}, Position{2, 2}, 4},
	{Position{0, 0}, Position{-2, -2}, 4},
	{Position{-1, -1}, Position{-2, -2}, 2},
	{Position{-1, 1}, Position{2, -2}, 6},
}


func TestDistance(t *testing.T) {
	for _, pair := range tests {
		v := Distance(pair.position1, pair.position2)
		if pair.distance != v {
			t.Error(
				"For", pair.position1.X, pair.position1.Y, " -> ",pair.position2.X, pair.position2.Y,
				"expected", pair.distance,
				"got", v,
			)
		}
	}
}