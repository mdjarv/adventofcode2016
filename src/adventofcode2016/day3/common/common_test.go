package common

import "testing"

type testpair struct {
	values   string
	possible bool
}

var testsPart1 = []testpair{
	{"5 10 25", false },
	{"10 5 25", false },
	{"25 10 5", false },
	{"10 10 10", true },
	{"10 5 6", true },
}

func TestValidTriangle(t *testing.T) {
	for _, pair := range testsPart1 {
		v := IsTrianglePossible(pair.values)
		if v != pair.possible {
			t.Error(
				"For", pair.values,
				"expected", pair.possible,
				"got", v,
			)
		}
	}
}