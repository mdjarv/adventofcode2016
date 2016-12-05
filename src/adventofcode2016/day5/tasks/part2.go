package tasks

import (
	"adventofcode2016/day5/common"
	"fmt"
	"strings"
)

func decryptPart2Password(door string) string {
	password := []byte{
		'_', '_', '_', '_',
		'_', '_', '_', '_',
	}
	offset := 0

	for strings.Contains(string(password), "_") {
		foundAt, hash := common.FindHash(offset, door)
		result := fmt.Sprintf("%x", hash)
		offset = foundAt + 1
		pos := result[5] - '0'
		if pos < 8 && password[pos] == '_' {
			password[pos] = result[6]
			//fmt.Printf("Found hash %x at %d, placing %c at position %d\n", hash, foundAt, result[6], pos)
			//fmt.Println("At offset", foundAt, "current password:", string(password))
		}
	}
	return string(password[:])
}

func Part2DecryptPassword() string {
	return decryptPart2Password("ojvtpuvg")
}