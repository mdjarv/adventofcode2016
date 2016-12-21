package tasks

import (
	"testing"
)

type testpair struct {
	input  input
	output interface{}
}

type input struct {
	state  string
	length int
}

var testDragonCurve = []testpair{
	{input{"1", 0}, "100"},
	{input{"0", 0}, "001"},
	{input{"11111", 0}, "11111000000"},
	{input{"111100001010", 0}, "1111000010100101011110000"},
}

var testChecksum = []testpair{
	{input{"110010110100", 12}, "100"},
}

var testSizedDragonCurve = []testpair{
	{input{"10000", 20}, "10000011110010000111"},
}

var testDataPart1 = []testpair{
	{input{"10000", 20}, "01100"},
}

var testDataPart2 = []testpair{
	{input{"10000", 0}, "01100"},
}

func TestStringToDragonCurve(t *testing.T) {
	for _, pair := range testDragonCurve {
		result := dragonCurve(pair.input.state)
		if result != pair.output {
			t.Error(
				"For input", pair.input.state,
				"expected", pair.output,
				"got", result,
			)
		}
	}
}

func TestSizedDragonCurve(t *testing.T) {
	for _, pair := range testSizedDragonCurve {
		result := sizedDragonCurve(pair.input.state, pair.input.length)
		if result != pair.output {
			t.Error(
				"For input", pair.input.state,
				"expected", pair.output,
				"got", result,
			)
		}
	}
}

func TestChecksum(t *testing.T) {
	for _, pair := range testChecksum {
		result := checksum(pair.input.state)
		if result != pair.output {
			t.Error(
				"For input", pair.input.state,
				"expected", pair.output,
				"got", result,
			)
		}
	}
}

func TestDay16Part1(t *testing.T) {
	for _, pair := range testDataPart1 {
		result := Day16(pair.input.state, pair.input.length)
		if result != pair.output {
			t.Error(
				"For input", pair.input,
				"expected", pair.output,
				"got", result,
			)
		}
	}
}
