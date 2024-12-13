package day13

import (
	"math"
	"strings"

	"misode.dev/aoc-2024/utils"
)

func Solve() (int, int) {
	timer := utils.StartDay(13)

	lines := utils.ReadInput("in.txt")
	groups := utils.SplitLinesOnEmpty(lines)

	part1 := 0
	for _, group := range groups {
		aParts := strings.Split(group[0][len("Button A: "):], ", ")
		ax, ay := utils.StrToInt(aParts[0][2:]), utils.StrToInt(aParts[1][2:])
		bParts := strings.Split(group[1][len("Button B: "):], ", ")
		bx, by := utils.StrToInt(bParts[0][2:]), utils.StrToInt(bParts[1][2:])
		pParts := strings.Split(group[2][len("Prize: "):], ", ")
		px, py := utils.StrToInt(pParts[0][2:]), utils.StrToInt(pParts[1][2:])

		minTokens := math.MaxInt
		possible := false

		for ac := 0; ac <= 100; ac++ {
			rx := (px - (ac * ax))
			if rx%bx == 0 {
				bc := rx / bx
				if ac*ay+bc*by == py {
					tokens := 3*ac + 1*bc
					possible = true
					if tokens < minTokens {
						minTokens = tokens
					}
				}
			}
		}

		if possible {
			part1 += minTokens
		}
	}
	timer.Part1(part1)

	part2 := 0
	timer.Part2(part2)

	return part1, part2
}
