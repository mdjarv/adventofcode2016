package tasks

import "testing"

type testpair struct {
	input string
	output string
}

var testDataPart1 = []testpair{
	{"ADVENT", "ADVENT"},
	{"A(1x5)BC", "ABBBBBC"},
	{"(3x3)XYZ", "XYZXYZXYZ"},
	{"A(2x2)BCD(2x2)EFG", "ABCBCDEFEFG"},
	{"(6x1)(1x3)A", "(1x3)A"},
	{"X(8x2)(3x3)ABCY", "X(3x3)ABC(3x3)ABCY"},
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
