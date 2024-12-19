package day17

import (
	"strings"

	"misode.dev/aoc-2024/utils"
)

func Solve() (string, int) {
	timer := utils.StartTimer()

	lines := utils.ReadInput("in.txt")
	groups := utils.SplitLinesOnEmpty(lines)
	a := utils.SplitInts(lines[0], " ")[0]
	b := utils.SplitInts(lines[1], " ")[0]
	c := utils.SplitInts(lines[2], " ")[0]
	program := utils.SplitInts(strings.Split(groups[1][0], " ")[1], ",")

	out := []int{}
	for p := 0; p+1 < len(program); p += 2 {
		op, literal := program[p], program[p+1]

		combo := literal
		if literal == 4 {
			combo = a
		} else if literal == 5 {
			combo = b
		} else if literal == 6 {
			combo = c
		}

		switch op {
		case 0: // adv
			a = a / (1 << combo)
		case 1: // bxl
			b = b ^ literal
		case 2: // bst
			b = combo % 8
		case 3: // jnz
			if a != 0 {
				p = literal - 2
			}
		case 4: // bxc
			b = b ^ c
		case 5: // out
			out = append(out, combo%8)
		case 6: // bdv
			b = a / (1 << combo)
		case 7: // cdv
			c = a / (1 << combo)
		}
	}
	part1 := utils.JoinInts(out, ",")
	timer.Part(part1)

	part2 := 0
	timer.Part(part2)

	return part1, part2
}
