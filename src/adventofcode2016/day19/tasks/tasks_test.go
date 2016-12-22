package tasks

import (
	"testing"
)

type testpair struct {
	input  int
	output interface{}
}

var testDataPart1 = []testpair{
	{5, 3},
}

var testDataPart2 = []testpair{
	{5, 2},
}

func TestDay19Part1(t *testing.T) {
	for _, pair := range testDataPart1 {
		result := Day19Part1(pair.input)
		if result != pair.output {
			t.Error(
				"For Part1",
				"expected", pair.output,
				"got", result,
			)
		}
	}
}

func TestDay19Part2(t *testing.T) {
	for _, pair := range testDataPart2 {
		result := Day19Part2(pair.input)
		if result != pair.output {
			t.Error(
				"For Part2",
				"expected", pair.output,
				"got", result,
			)
		}
	}
}

func TestTradePresents(t *testing.T) {
	input := []Elf {
		Elf{position:1, presents:1},
		Elf{position:2, presents:1},
	}
	expected := []Elf {
		Elf{position:1, presents:2},
		Elf{position:2, presents:0},
	}

	output := tradePresents(input)

	for i := range expected {
		if output[i] != expected[i] {
			t.Errorf("Expected %+v got %+v\n", expected[i], output[i])
		}
	}
}
