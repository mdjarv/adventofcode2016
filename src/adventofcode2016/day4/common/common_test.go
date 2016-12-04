package common

import (
	"testing"
	"adventofcode2016/day4/tasks"
)

type testpair struct {
	value  string
	result interface{}
}

var testsRoomFromString = []testpair{
	{"aaaaa-bbb-z-y-x-123[abxyz]", Room{"aaaaa-bbb-z-y-x", 123, "abxyz"} },
	{"a-b-c-d-e-f-g-h-987[abcde]", Room{"a-b-c-d-e-f-g-h", 987, "abcde"} },
	{"not-a-real-room-404[oarel]", Room{"not-a-real-room", 404, "oarel"} },
	{"totally-real-room-200[decoy]", Room{"totally-real-room", 200, "decoy"} },
}

var testsGenerateChecksum = []testpair{
	{"aaaaa-bbb-z-y-x-123", "abxyz" },
	{"a-b-c-d-e-f-g-h-987", "abcde" },
	{"not-a-real-room-404", "oarel" },
}

var testsPart1 = []testpair{
	{"aaaaa-bbb-z-y-x-123[abxyz]", true },
	{"a-b-c-d-e-f-g-h-987[abcde]", true },
	{"not-a-real-room-404[oarel]", true },
	{"totally-real-room-200[decoy]", false },
}

func TestRoomFromString(t *testing.T) {
	for _, pair := range testsRoomFromString {
		v := RoomFromString(pair.value)
		if v != pair.result {
			t.Error(
				"For", pair.value,
				"expected", pair.result,
				"got", v,
			)
		}
	}
}

func TestGenerateChecksum(t *testing.T) {
	for _, pair := range testsGenerateChecksum {
		v := generateChecksum(pair.value)
		if v != pair.result {
			t.Error(
				"For", pair.value,
				"expected", pair.result,
				"got", v,
			)
		}
	}
}

func TestIsRoomReal(t *testing.T) {
	for _, pair := range testsPart1 {
		v := IsRoomReal(RoomFromString(pair.value))
		if v != pair.result {
			t.Error(
				"For", pair.value,
				"expected", pair.result,
				"got", v,
			)
		}
	}
}