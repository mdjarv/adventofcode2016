package tasks

import "testing"

type testpair struct {
	input  string
	output bool
}

var testDataPart1 = []testpair{
	{"abba[mnop]qrst", true},
	{"abcd[bddb]xyyx", false},
	{"aaaa[qwer]tyui", false},
	{"ioxxoj[asdfgh]zxcvbn", true},
}

func TestDay7Part1(t *testing.T) {
	for _, pair := range testDataPart1 {
		_, result := CheckIPv7Address(pair.input)
		if result != pair.output {
			t.Error(
				"For", pair.input,
				"expected", pair.output,
				"got", result,
			)
		}
	}
}
