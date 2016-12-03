package tasks

import "testing"

type testpair struct {
	values   string
	distance int
}

var testsPart1 = []testpair{
	{"R2, L3", 5 },
	{"R2, R2, R2", 2 },
	{"R5, L5, R5, R3", 12 },
}

var testsPart2 = []testpair{
	{"R8, R4, R4, R8", 4 },
}

func TestPart1(t *testing.T) {
	for _, pair := range testsPart1 {
		v := GetPart1Distance(pair.values)
		if v != pair.distance {
			t.Error(
				"For", pair.values,
				"expected", pair.distance,
				"got", v,
			)
		}
	}
}

func TestPart2(t *testing.T) {
	for _, pair := range testsPart2 {
		v := GetPart2Distance(pair.values)
		if v != pair.distance {
			t.Error(
				"For", pair.values,
				"expected", pair.distance,
				"got", v,
			)
		}
	}
}