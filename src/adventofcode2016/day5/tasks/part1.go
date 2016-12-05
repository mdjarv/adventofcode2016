package tasks

import (
	"adventofcode2016/day5/common"
	"bytes"
	"fmt"
)

func decryptPart1Password(door string) string {
	var password bytes.Buffer
	offset := 0
	for password.Len() < 8 {
		foundAt, hash := common.FindHash(offset, door)
		result := fmt.Sprintf("%x", hash)
		offset = foundAt + 1
		password.WriteString(string(result[5]))
	}
	return password.String()
}

func Part1DecryptPassword() string {
	return decryptPart1Password("ojvtpuvg")
}