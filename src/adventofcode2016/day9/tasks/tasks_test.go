package tasks

import (
	"testing"
	"fmt"
)

type testpair struct {
	input string
	output interface{}
}

var testDataPart1 = []testpair{
	{"ADVENT", "ADVENT"},
	{"A(1x5)BC", "ABBBBBC"},
	{"(3x3)XYZ", "XYZXYZXYZ"},
	{"A(2x2)BCD(2x2)EFG", "ABCBCDEFEFG"},
	{"(6x1)(1x3)A", "(1x3)A"},
	{"X(8x2)(3x3)ABCY", "X(3x3)ABC(3x3)ABCY"},
}

var testDataPart2 = []testpair{
	{"ADVENT", 6},
	{"A(1x5)BC", 7},
	{"(3x3)XYZ", 9},
	{"A(2x2)BCD(2x2)EFG", 11},
	{"(6x1)(1x3)A", 3},
	{"X(8x2)(3x3)ABCY", 20},
	{"(27x12)(20x12)(13x14)(7x10)(1x12)A", 241920},
	{"(25x3)(3x3)ABC(2x3)XY(5x2)PQRSTX(18x9)(3x2)TWO(5x7)SEVEN", 445},
}

func TestDay8Part1(t *testing.T) {
	for _, pair := range testDataPart1 {
		result := DecompressString(pair.input)
		if result != pair.output {
			t.Error(
				"For", pair.input,
				"expected", pair.output,
				"got", result,
			)
		}
	}
}

func TestDay8Part2(t *testing.T) {
	for _, pair := range testDataPart2 {
		fmt.Println()
		result := Day9Part2(pair.input)
		if result != pair.output {
			t.Error(
				"For", pair.input,
				"expected", pair.output,
				"got", result,
			)
		}
	}
}