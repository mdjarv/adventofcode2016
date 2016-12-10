package tasks

import (
	"regexp"
	"strconv"
	"bytes"
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

func findAllMarkers(line string) ([]Marker, int) {
	markers := []Marker{}
	remainder := bytes.Buffer{}

	marker, ok := findMarker(line, 0)
	if !ok {
		// No markers
		remainder.WriteString(line)
	}
	for ok {
		markers = append(markers, marker)
		marker, ok = findMarker(line, marker.end + marker.characters)
	}

	for i, marker := range markers {
		if i == 0 && markers[i].start > 0 {
			remainder.WriteString(line[:marker.start])
		}
		if len(markers) > 1 && i > 0 {
			remainder.WriteString(line[markers[i - 1].end + markers[i - 1].characters:markers[i].start])
		}
		if i == len(markers) - 1 {
			s := marker.end + marker.characters
			if s < len([]rune(line)) {
				remainder.WriteString(line[s:])
			}
		}
	}

	return markers, remainder.Len()
}

func getV2Size(line string) int {
	markers, size := findAllMarkers(line)
	if len(markers) == 0 {
		return size
	} else {
		for _, marker := range markers {
			size += getV2Size(marker.str) * marker.repeat
		}
	}
	return size
}

func Day9Part2(line string) int {
	return getV2Size(line)
}
