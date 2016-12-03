package tasks

import (
	"adventofcode2016/day1/common"
)

func GetPart1Distance(command string) int {
	path := common.GetPath(command)
	start := path[0]
	end := path[len(path) - 1]
	return common.Distance(start, end)
}