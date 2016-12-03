package tasks

import (
	"testing"
	"adventofcode2016/day2/common"
)

func TestPart2(t *testing.T) {
	data := []string{"ULL", "RRDDD", "LURDL", "UUUUD"}
	expected := "5DB3"

	code := common.GetCode(data, Part2Keypad, Part2InitialPosition)
	if code != expected {
		t.Error("Expected", expected, "got", code)
	}
}