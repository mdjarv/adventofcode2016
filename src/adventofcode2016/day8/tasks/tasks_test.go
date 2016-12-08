package tasks

import "testing"

type testpair struct {
	input  []string
	output int
}

var testDataPart1 = []testpair{
	{
		[]string{
			"rect 3x2",
			"rotate column x=1 by 1",
			"rotate row y=0 by 4",
			"rotate column x=1 by 1",
		}, 6,
	},
}

func TestDay8Part1(t *testing.T) {
	for _, pair := range testDataPart1 {
		result := Day8Part1(pair.input, 7, 3)
		if result != pair.output {
			t.Error(
				"For", pair.input,
				"expected", pair.output,
				"got", result,
			)
		}
	}
}
