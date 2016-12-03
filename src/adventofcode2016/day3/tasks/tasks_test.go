package tasks

import "testing"

type testpair struct {
	values   []string
	possible int
}

var testsPart1 = []testpair{
	{
		[]string{
			"5 10 25",
			"10 5 25",
			"25 10 5",
			"10 10 10",
			"10 5 6",
		}, 2,
	},
}

func TestPart1(t *testing.T) {
	for _, pair := range testsPart1 {
		v := Part1ValidTriangleCount(pair.values)
		if v != pair.possible {
			t.Error(
				"For", pair.values,
				"expected", pair.possible,
				"got", v,
			)
		}
	}
}
