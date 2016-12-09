package tasks

import (
	"regexp"
	"strconv"
	"bytes"
	"fmt"
)

type Marker struct {
	characters int
	repeat     int
	start      int
	end        int
	str        string
}

func findMarker(line string, offset int) (Marker, bool) {
	marker := Marker{}
	markerRegexp := regexp.MustCompile(`\((?P<count>\d+)x(?P<times>\d+)\)`)

	markers := markerRegexp.FindAllStringIndex(line, -1)
	if len(markers) > 0 {
		for _, found := range markers {
			start := found[0]
			end := found[1]
			if start >= offset {
				match := markerRegexp.FindStringSubmatch(line[start:end])
				marker.characters, _ = strconv.Atoi(match[1])
				marker.repeat, _ = strconv.Atoi(match[2])
				marker.start = start
				marker.end = end
				marker.str = line[marker.end:(marker.end + marker.characters)]

				return marker, true
			}
		}

	}
	return marker, false
}

func DecompressString(line string) string {
	tries := 0
	marker, ok := findMarker(line, 0)
	for ok && tries < 10 {
		tries++
		var buffer bytes.Buffer
		prefix := line[:marker.start]
		repeat := line[marker.end:marker.end + marker.characters]
		remainder := line[marker.end + marker.characters:]

		buffer.WriteString(prefix)
		for i := 0; i < marker.repeat; i++ {
			buffer.WriteString(repeat)
		}
		buffer.WriteString(remainder)

		line = buffer.String()
		marker, ok = findMarker(line, marker.start + (marker.characters * marker.repeat))
	}
	return line
}

func Day9Part1(line string) int {
	return len([]rune(DecompressString(line)))
}

func getSizeOfData(line string, offset int) (int, string) {
	fmt.Println("processing", line, "from offset", offset)
	size := 0
	remainder := ""

	marker, ok := findMarker(line, offset)
	if !ok {
		size += len([]rune(line))
		fmt.Println("  no markers, returning size", size)
		return size, remainder
	}
	for ok {
		prefix := line[offset:marker.start]
		prefixSize := len([]rune(prefix))
		size += prefixSize
		fmt.Println("  prefix", prefix, "size", prefixSize, "marker inner content", marker.str)
		fmt.Println("  processing marker", marker.str, "*", marker.repeat)

		innerMarker, hasInnerMarker := findMarker(marker.str, 0)
		if hasInnerMarker {
			fmt.Println("  has inner marker", innerMarker.str)
			innerSize, _ := getSizeOfData(innerMarker.str, 0)
			fmt.Println("  inner marker returned with size", innerSize, "multiplying with", innerMarker.repeat)
			size += innerSize * innerMarker.repeat * marker.repeat
		} else {
			currentSize := len([]rune(marker.str)) * marker.repeat
			fmt.Println("  no inner marker, appending", currentSize)
			size += currentSize
		}

		remainder = line[marker.end + marker.characters:]
		offset = marker.end + marker.characters
		marker, ok = findMarker(line, marker.end + marker.characters)
	}

	fmt.Println("  returning", size)
	return size, remainder
}

func Day9Part2(line string) int {
	size, remainder := getSizeOfData(line, 0)
	size += len([]rune(remainder))
	fmt.Println("FINAL SIZE", size)

	return size
}
