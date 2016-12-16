package tasks

import (
	"testing"
)

type testpair struct {
	input  []string
	output interface{}
}

var inputData1 = []string{
	"cpy 41 a",
	"inc a",
	"inc a",
	"dec a",
	"jnz a 2",
	"dec a",
}

var inputData2 = []string{
	"cpy 41 a",
	"inc a",
	"inc a",
	"dec a",
	"jnz 1 2",
	"dec a",
}

var testDataPart1 = []testpair{
	{inputData1, 42},
	{inputData2, 42},
}

func TestDay12Part1(t *testing.T) {
	for _, pair := range testDataPart1 {
		result := Day12Part1(pair.input)
		if result != pair.output {
			t.Error(
				"For Part1",
				"expected", pair.output,
				"got", result,
			)
		}
	}
}
