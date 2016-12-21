package tasks

import (
	"bytes"
	"fmt"
)

func sizedDragonCurve(a string, length int) string {
	r := dragonCurve(a)
	for len(r) < length {
		r = dragonCurve(r)
	}
	return r[:length]
}

func dragonCurve(a string) string {
	b := reverse(a)
	b = invert(b)
	return fmt.Sprintf("%s0%s", a, b)
}

func invert(in string) string {
	buffer := bytes.Buffer{}
	for _, r := range in {
		if r == '1' {
			buffer.WriteRune('0')
		} else {
			buffer.WriteRune('1')
		}
	}
	return buffer.String()
}

func checksum(in string) string {
	c := generateChecksum(in)
	for (len(c) % 2) == 0 {
		c = generateChecksum(c)
	}
	return c
}

func generateChecksum(in string) string {
	buffer := bytes.Buffer{}
	for i := 0; i < len(in); i += 2 {
		if in[i] == in[i + 1] {
			buffer.WriteRune('1')
		} else {
			buffer.WriteRune('0')
		}
	}
	return buffer.String()
}

func reverse(in string) string {
	buffer := bytes.Buffer{}
	for i := len(in) - 1; i >= 0; i-- {
		buffer.WriteRune(rune(in[i]))
	}
	return buffer.String()
}

func Day16(input string, length int) string {
	return checksum(sizedDragonCurve(input, length))
}
