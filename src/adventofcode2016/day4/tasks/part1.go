package tasks

import "adventofcode2016/day4/common"

func Part1SumOfValidSectors(lines []string) int {
	sectorSum := 0
	for _, room := range common.GetValidRooms(lines) {
		sectorSum += room.Sector
	}

	return sectorSum
}