package tasks

import (
	"adventofcode2016/day1/common"
)

func GetPart2Distance(command string) int {
	path := common.GetPath(command)
	start := path[0]
	for i, pos1 := range path[1:] {
		for _, pos2 := range path[:i] {
			if pos1.X == pos2.X && pos1.Y == pos2.Y {
				return common.Distance(start, pos1)
			}
		}
	}
	return -1
}