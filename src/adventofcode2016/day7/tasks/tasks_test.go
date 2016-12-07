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
	{"asdffe[asdfgh]zxcvbn[asdbaab]fesfaabba", false},
	{"asdfge[asdfgh]zxcvbn[asdbeab]fesfaabba", true},
}

var testDataPart2 = []testpair{
	{"aba[bab]xyz", true},
	{"xyx[xyx]xyx", false},
	{"aaa[kek]eke", true},
	{"zazbz[bzb]cdb", true},
}

func TestDay7Part1(t *testing.T) {
	for _, pair := range testDataPart1 {
		result := IPv7AddressTLSSupport(pair.input)
		if result != pair.output {
			t.Error(
				"For", pair.input,
				"expected", pair.output,
				"got", result,
			)
		}
	}
}

func TestDay7Part2(t *testing.T) {
	for _, pair := range testDataPart2 {
		result := IPv7AddressSSLSupport(pair.input)
		if result != pair.output {
			t.Error(
				"For", pair.input,
				"expected", pair.output,
				"got", result,
			)
		}
	}
}
