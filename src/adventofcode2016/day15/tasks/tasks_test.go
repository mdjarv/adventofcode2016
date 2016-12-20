package tasks

import (
	"testing"
)

type testpair struct {
	input  []string
	output interface{}
}

var testDataPart1 = []testpair{
	{
		[]string{
			"Disc #1 has 5 positions; at time=0, it is at position 4.",
			"Disc #2 has 2 positions; at time=0, it is at position 1.",
		}, 5},
}

func TestDay15Part1(t *testing.T) {
	for _, pair := range testDataPart1 {
		result := Day15Part1(pair.input)
		if result != pair.output {
			t.Error(
				"For Part1",
				"expected", pair.output,
				"got", result,
			)
		}
	}
}

