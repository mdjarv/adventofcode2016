package tasks

import (
	"fmt"
	"regexp"
	"strconv"
	"bytes"
)

type Marker struct {
	count int
	times int
	start int
	end int
}

func findMarker(line string) (Marker, bool) {
	marker := Marker{}
	return marker, markerFound
}

func DecompressString(line string) string {
	fmt.Println("decompressing", line)
	markerRegexp := regexp.MustCompile(`\((?P<count>\d+)x(?P<times>\d+)\)`)

	marker := markerRegexp.FindStringIndex(line)
	if len(marker) > 0 {
		match := markerRegexp.FindStringSubmatch(line)
		fmt.Println("match",match)
		fmt.Println("marker",marker)
		count, _ := strconv.Atoi(match[1])
		//times := match[2]
		premark := line[:marker[0]]
		repeat := line[marker[1]:marker[1]+count]
		remainder := line[marker[1]+count:]
		fmt.Println(premark, repeat, remainder)
		var buffer bytes.Buffer
		buffer.WriteString(premark)
		for i := 0; i < match[2]; i++ {
			buffer.WriteString(repeat)
		}
		buffer.WriteString(remainder)
		fmt.Println(buffer.String())
		marker = markerRegexp.FindStringIndex(line)
	}
	return line
}

func Day9Part1(line string) int {
	return len([]rune(DecompressString(line)))
}
