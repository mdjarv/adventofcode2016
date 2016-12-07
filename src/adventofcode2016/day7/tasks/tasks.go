package tasks

import "fmt"

type IPv7Address struct {
	part1 string
	part2 string
	part3 string
}

func CheckIPv7Address(input string) (IPv7Address, bool) {
	address := IPv7Address{}

	return address, false
}

func Day7Part1(lines []string) int {
	for _, line := range lines {
		address, valid := CheckIPv7Address(line)
		fmt.Print(address)
		if valid {
			fmt.Println("valid")
		} else {
			fmt.Println("not valid")
		}
	}

	return 0
}