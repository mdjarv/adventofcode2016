package tasks

import (
	"strings"
	"strconv"
)

func processCommands(register map[string]int, lines []string) map[string]int {
	i := 0
	//fmt.Printf("AAAA BBBB CCCC DDDD   CMD\n")
	for i < len(lines) {
		//fmt.Printf("%4d %4d %4d %4d   %s\n", register["a"], register["b"], register["c"], register["d"], lines[i])
		args := strings.Split(lines[i], " ")
		switch(args[0]) {
		case "cpy":
			val, err := strconv.Atoi(args[1])
			if err != nil {
				register[args[2]] = register[args[1]]
			} else {
				register[args[2]] = val
			}
		case "inc":
			register[args[1]]++
		case "dec":
			register[args[1]]--
		case "jnz":
			val, _ := strconv.Atoi(args[2])
			check, err := strconv.Atoi(args[1])
			if err != nil {
				check = register[args[1]]
			}

			if check != 0 {
				i += val
				continue
			}
		}
		i++
	}
	return register
}

func Day12Part1(lines []string) int {
	register := map[string]int{}
	register["a"] = 0
	register["b"] = 0
	register["c"] = 0
	register["d"] = 0

	register = processCommands(register, lines)

	return register["a"]
}

func Day12Part2(lines []string) int {
	register := map[string]int{}
	register["a"] = 0
	register["b"] = 0
	register["c"] = 1
	register["d"] = 0

	register = processCommands(register, lines)

	return register["a"]
}
