package tasks

import (
	"testing"
)

type testpair struct {
	input  int
	output interface{}
}

var testDataPart1 = []testpair{
	{10, 11},
}

func TestDay13Part1(t *testing.T) {
	for _, pair := range testDataPart1 {
		result := Day13Part1(pair.input)
		if result != pair.output {
			t.Error(
				"For Part1",
				"expected", pair.output,
				"got", result,
			)
		}
	}
}
