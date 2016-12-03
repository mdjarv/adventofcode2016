package tasks

import (
	"testing"
	"adventofcode2016/day2/common"
)

func TestPart1(t *testing.T) {
	data := []string{"ULL", "RRDDD", "LURDL", "UUUUD"}
	expected := "1985"

	code := common.GetCode(data, Part1Keypad, Part1InitialPosition)
	if code != expected {
		t.Error("Expected", expected, "got", code)
	}
}