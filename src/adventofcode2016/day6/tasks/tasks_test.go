package tasks

import "testing"

type testpair struct {
	input []string
	output string
}

var inputTestData = []string {
	"eedadn",
	"drvtee",
	"eandsr",
	"raavrd",
	"atevrs",
	"tsrnev",
	"sdttsa",
	"rasrtv",
	"nssdts",
	"ntnada",
	"svetve",
	"tesnvt",
	"vntsnd",
	"vrdear",
	"dvrsen",
	"enarar",
}

var testDataPart1 = []testpair {
	{inputTestData, "easter"},
}

var testDataPart2 = []testpair {
	{inputTestData, "advent"},
}

func TestPart1SignalsAndNoise(t *testing.T) {
	for _, pair := range testDataPart1 {
		result := Part1SignalsAndNoise(pair.input)
		if result != pair.output {
			t.Error(
				"For", pair.input,
				"expected", pair.output,
				"got", result,
			)
		}
	}
}

func TestPart2SignalsAndNoise(t *testing.T) {
	for _, pair := range testDataPart2 {
		result := Part2SignalsAndNoise(pair.input)
		if result != pair.output {
			t.Error(
				"For", pair.input,
				"expected", pair.output,
				"got", result,
			)
		}
	}
}