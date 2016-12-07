package tasks

import (
	"regexp"
	"unicode/utf8"
	"strings"
)

type IPv7Address struct {
	hypernets []string
	supernets []string
}

func stringToIPv7Address(s string) IPv7Address {
	reHyper := regexp.MustCompile(`\[(?P<hypernet>\w+)]`)

	hypernets := reHyper.FindAllString(s, -1)
	supernets := strings.Split(reHyper.ReplaceAllString(s, "|"), "|")

	return IPv7Address{hypernets, supernets}
}

func IPv7AddressTLSSupport(input string) bool {
	address := stringToIPv7Address(input)
	for _, hypernet := range address.hypernets {
		if isABBA(hypernet) {
			return false
		}
	}

	for _, part := range address.supernets {
		if isABBA(part) {
			return true
		}
	}

	return false
}

func isABBA(s string) bool {
	length := utf8.RuneCountInString(s)
	for i := range s {
		if i > (length - 4) {
			break
		}
		if (s[i] != s[i + 1] && s[i] == s[i + 3] && s[i + 1] == s[i + 2]) {
			return true
		}
	}
	return false
}

func getABAList(s string) []string {
	length := utf8.RuneCountInString(s)
	aba := []string{}
	for i := range s {
		if i > (length - 3) {
			break
		}
		if (s[i] == s[i + 2] && s[i] != s[i+1]) {
			aba = append(aba, s[i:i + 3])
		}
	}
	return aba
}

func IPv7AddressSSLSupport(input string) bool {
	address := stringToIPv7Address(input)
	for _, supernet := range address.supernets {
		abaList := getABAList(supernet)
		for _, aba := range abaList {
			bab := string([]byte{aba[1], aba[0], aba[1]})
			for _, hypernet := range address.hypernets {
				if strings.Contains(hypernet, bab) {
					return true
				}
			}
		}
	}
	return false
}

func Day7Part1(lines []string) int {
	validAddresses := 0
	for _, line := range lines {
		if IPv7AddressTLSSupport(line) {
			validAddresses++
		}
	}

	return validAddresses
}

func Day7Part2(lines []string) int {
	validAddresses := 0
	for _, line := range lines {
		if IPv7AddressSSLSupport(line) {
			validAddresses++
		}
	}

	return validAddresses
}