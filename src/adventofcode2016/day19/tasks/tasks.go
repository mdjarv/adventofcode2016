package tasks

import (
	"fmt"
	"time"
)

type Elf struct {
	position int
	presents int
}

func elvesLeft(elves []Elf) int {
	elvesWithPresents := 0
	for _, elf := range elves {
		if elf.presents > 0 {
			elvesWithPresents++
		}
	}
	return elvesWithPresents
}

func nextPositionWithPresents(elves []Elf, position int) int {
	for i := 1; i <= len(elves); i++ {
		p := (i + position) % len(elves)
		if (elves[p].presents > 0) {
			return p
		}
	}
	return position
}

func tradePresents(elves []Elf) []Elf {
	for i, elf := range elves {
		if elf.presents > 0 {
			next := nextPositionWithPresents(elves, i)
			if next != i {
				nextElf := elves[next]
				elf.presents += nextElf.presents
				nextElf.presents = 0
				elves[i] = elf
				elves[next] = nextElf
			}
		}
	}
	return elves
}

func populateElves(num int) []Elf {
	elves := []Elf{}
	for i := 0; i < num; i++ {
		elves = append(elves, Elf{i + 1, 1})
	}
	return elves
}

func Day19Part1(numElves int) int {
	elves := populateElves(numElves)
	for elvesLeft(elves) > 1 {
		elves = tradePresents(elves)
	}

	for _, elf := range elves {
		if elf.presents > 0 {
			return elf.position
		}
	}
	return -1
}

var currentTime = time.Now().UnixNano()

func getOpposite(elves []Elf, trader int) int {
	playing := 0
	for _, elf := range elves {
		if (elf.presents > 0) {
			playing++
		}
	}
	opposite := (playing + 1) / 2 // round up
	if playing % 1000 == 0 {
		t := time.Now().UnixNano()
		fmt.Println((t-currentTime)/1000000, "ms passed,", playing, "still playing")
		currentTime = t
	}

	current := 0
	for _, elf := range elves {
		if (elf.presents > 0) {
			current++
			if current >= opposite {
				return elf.position
			}
		}
	}

	return -1
}

func findElf(elves []Elf, pos int) (int, Elf) {
	fmt.Println("Finding elf", pos)
	//for i, elf := range elves {
	//	if elf.position == pos {
	//		return i, elf
	//	}
	//}
	//return -1, Elf{0, 0}
	return pos-1, elves[pos]
}

func tradeAcross(elves []Elf, traderPos int) []Elf {
	oppositePos := getOpposite(elves, traderPos)

	traderIndex, trader := findElf(elves, traderPos)
	oppositeIndex, opposite := findElf(elves, oppositePos)

	trader.presents += opposite.presents
	opposite.presents = 0

	elves[traderIndex] = trader
	elves[oppositeIndex] = opposite

	return elves
}

func nextTrader(elves []Elf, trader int) int {
	for i := 1; i < len(elves); i++ {
		pos := (trader + i) % len(elves)
		_, elf := findElf(elves, pos)
		if elf.presents > 0 {
			//fmt.Println("next trader", elf)
			return pos
		}
	}

	return trader
}

func Day19Part2(numElves int) int {
	elves := populateElves(numElves)
	trader := 1
	lastTrade := -1

	for trader != lastTrade {
		elves = tradeAcross(elves, trader)
		lastTrade = trader
		trader = nextTrader(elves, trader)
	}
	return elves[trader].position
}
