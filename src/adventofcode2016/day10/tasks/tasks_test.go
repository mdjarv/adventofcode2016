package tasks

import (
	"testing"
)

type testpair struct {
	input  []string
	output interface{}
}

var testDataPart1 = []testpair{
	{[]string{
		"value 5 goes to bot 2",
		"bot 2 gives low to bot 1 and high to bot 0",
		"value 3 goes to bot 1",
		"bot 1 gives low to output 1 and high to bot 0",
		"bot 0 gives low to output 2 and high to output 0",
		"value 2 goes to bot 2",
	}, 2},
}

func TestDay10Part1(t *testing.T) {
	for _, pair := range testDataPart1 {
		result := Day10Part1(pair.input, 5, 2)
		if result != pair.output {
			t.Error(
				"For", pair.input,
				"expected", pair.output,
				"got", result,
			)
		}
	}
}
