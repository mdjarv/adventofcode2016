package common

import (
	"crypto/md5"
	"fmt"
	"bytes"
	"strconv"
	"strings"
)

func isCorrectHash(h [16]byte) bool {
	p := fmt.Sprintf("%x", h)[:5]
	return strings.HasPrefix(p, "00000")
}

func makeHashableBytes(offset int, door string) []byte {
	var buffer bytes.Buffer
	buffer.WriteString(door)
	buffer.WriteString(strconv.Itoa(offset))
	return buffer.Bytes()
}

func FindHash(offset int, door string) (int, [16]byte) {
	currentHash := md5.Sum(makeHashableBytes(offset, door))
	for isCorrectHash(currentHash) == false {
		offset++
		currentHash = md5.Sum(makeHashableBytes(offset, door))
	}
	return offset, currentHash
}
