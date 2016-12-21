package tasks

import (
	"bytes"
	"strings"
)

var safe = '.'
var trap = '^'

var trapLookup = map[string]rune{
	"...": safe,
	"..^": trap,
	".^.": safe,
	".^^": trap,
	"^..": trap,
	"^.^": safe,
	"^^.": trap,
	"^^^": safe,
}

func generateTiles(line string) string {
	buffer := bytes.Buffer{}

	for i := 0; i < len(line); i++ {
		cb := bytes.Buffer{}

		if i > 0 {
			cb.WriteByte(line[i - 1])
		} else {
			cb.WriteRune(safe)
		}

		cb.WriteRune(rune(line[i]))

		if i < len(line) - 1 {
			cb.WriteByte(line[i + 1])
		} else {
			cb.WriteRune(safe)
		}

		lookup := cb.String()

		buffer.WriteRune(trapLookup[lookup])
	}

	return buffer.String()
}



func Day18Part1(line string, count int) int {
	safeCount := strings.Count(line, ".")
	for i := 0; i < count-1; i++ {
		line = generateTiles(line)
		safeCount += strings.Count(line, ".")
	}

	return safeCount
}
