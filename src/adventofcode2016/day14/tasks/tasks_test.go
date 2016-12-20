package tasks

import (
	"testing"
)

type testpair struct {
	input  string
	output interface{}
}

var testDataPart1 = []testpair{
	{"abc", 22728},
}

var testDataPart2 = []testpair{
	{"abc", 22859},
}

//func TestDay14Part1(t *testing.T) {
//	for _, pair := range testDataPart1 {
//		result := Day14Part1(pair.input)
//		if result != pair.output {
//			t.Error(
//				"For Part1",
//				"expected", pair.output,
//				"got", result,
//			)
//		}
//	}
//}

func TestGenerateStretchHash(t *testing.T) {
	result := GenerateStretchHash("abc", 0)
	expected := "a107ff634856bb300138cac6568c0f24"

	if result != expected {
		t.Error("For GenerateStretchHash expected", expected, "got", result)
	}
}

func TestDay14Part2(t *testing.T) {
	for _, pair := range testDataPart2 {
		result := Day14Part2(pair.input)
		if result != pair.output {
			t.Error(
				"For Part2",
				"expected", pair.output,
				"got", result,
			)
		}
	}
}
