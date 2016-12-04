package tasks

import (
	"testing"
	"adventofcode2016/day4/common"
)

type testpair struct {
	value  interface{}
	result interface{}
}

var testsPart1 = []testpair{
	{
		[]string {
			"aaaaa-bbb-z-y-x-123[abxyz]",
			"a-b-c-d-e-f-g-h-987[abcde]",
			"not-a-real-room-404[oarel]",
			"totally-real-room-200[decoy]",
		},
		1514,
	},
}


var testsDecryptRoom = []testpair{
	{common.Room{"qzmt-zixmtkozy-ivhz", 343, "aaaaa"}, common.Room{"very encrypted name", 343, "aaaaa"} },
}

func TestDecryptRoom(t *testing.T) {
	for _, pair := range testsDecryptRoom {
		v := DecryptRoom(pair.value.(common.Room))
		if v != pair.result {
			t.Error(
				"For", pair.value,
				"expected", pair.result,
				"got", v,
			)
		}
	}
}


func TestPart1(t *testing.T) {
	for _, pair := range testsPart1 {
		v := Part1SumOfValidSectors(pair.value.([]string))
		if v != pair.result {
			t.Error(
				"For", pair.value,
				"expected", pair.result,
				"got", v,
			)
		}
	}
}