package tasks

import (
	"testing"
)

type testpair struct {
	input  []string
	output interface{}
}

var inputData = []string {
	"The first floor contains a hydrogen-compatible microchip and a lithium-compatible microchip.",
	"The second floor contains a hydrogen generator.",
	"The third floor contains a lithium generator.",
	"The fourth floor contains nothing relevant.",
}

var testDataPart1 = []testpair{
	{inputData, 11},
}


func TestDay8Part1(t *testing.T) {
	for _, pair := range testDataPart1 {
		result := Day11Part1(pair.input)
		if result != pair.output {
			t.Error(
				"For Part1",
				"expected", pair.output,
				"got", result,
			)
		}
	}
}

