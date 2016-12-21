package tasks

import (
	"testing"
)

type testpair struct {
	input  string
	output interface{}
}

var testDataPart1 = []testpair{
	{"1", "100"},
	{"0", "001"},
	{"11111", "11111000000"},
	{"111100001010", "1111000010100101011110000"},
}

var testDataPart2 = []testpair{
}

func TestDay16Part1(t *testing.T) {
	for _, pair := range testDataPart1 {
		result := Day16Part1(pair.input)
		if result != pair.output {
			t.Error(
				"For input", pair.input,
				"expected", pair.output,
				"got", result,
			)
		}
	}
}

func TestDay16Part2(t *testing.T) {
	for _, pair := range testDataPart2 {
		result := Day16Part2(pair.input)
		if result != pair.output {
			t.Error(
				"For input", pair.input,
				"expected", pair.output,
				"got", result,
			)
		}
	}
}
