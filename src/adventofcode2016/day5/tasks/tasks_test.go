package tasks

import "testing"

type testpair struct {
	value  string
	result interface{}
}

var testPart1Data = []testpair{
	{"abc", "18f47a30"},
}

var testPart2Data = []testpair{
	{"abc", "05ace8e3"},
}

//func TestRoomFromString(t *testing.T) {
//	for _, pair := range testPart1Data {
//		v := decryptPart1Password(pair.value)
//		if v != pair.result {
//			t.Error(
//				"For", pair.value,
//				"expected", pair.result,
//				"got", v,
//			)
//		}
//	}
//}

func TestPart2DecryptPassword(t *testing.T) {
	for _, pair := range testPart2Data {
		v := decryptPart2Password(pair.value)
		if v != pair.result {
			t.Error(
				"For", pair.value,
				"expected", pair.result,
				"got", v,
			)
		}
	}
}