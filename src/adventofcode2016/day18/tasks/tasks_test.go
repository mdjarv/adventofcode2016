package tasks

import (
	"testing"
)

type tilespair struct {
	input  string
	output interface{}
}

var testGenerateTiles = []tilespair{
	{"..^^.", ".^^^^"},
	{".^^^^", "^^..^"},
	{".^^.^.^^^^", "^^^...^..^"},
	{"^^^...^..^", "^.^^.^.^^."},
	{"^.^^.^.^^.", "..^^...^^^"},
	{"..^^...^^^", ".^^^^.^^.^"},
	{".^^^^.^^.^", "^^..^.^^.."},
	{"^^..^.^^..", "^^^^..^^^."},
	{"^^^^..^^^.", "^..^^^^.^^"},
	{"^..^^^^.^^", ".^^^..^.^^"},
	{".^^^..^.^^", "^^.^^^..^^"},
}

func TestGenerateTiles(t *testing.T) {
	for _, pair := range testGenerateTiles {
		result := generateTiles(pair.input)
		if result != pair.output {
			t.Error(
				"Input", pair.input,
				"got", result,
				"expected", pair.output,
			)
		}
	}
}

func TestDay18Part1(t *testing.T) {
	result := Day18Part1(".^^.^.^^^^", 10)
	if result != 38 {
		t.Error(
			"got", result,
			"expected", 38,
		)
	}
}

