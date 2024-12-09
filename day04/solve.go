package day04

import (
	"misode.dev/aoc-2024/utils"
)

func Solve() (int, int) {
	timer := utils.StartDay(4)

	lines := utils.ReadInput("in.txt")
	G := utils.MakeGrid(lines)

	part1 := 0
	G.ForEach(func(r int, c int, _ rune) {
		if G.MatchSym(r, c, "XMAS") {
			part1 += 1
		}
		if G.MatchSym(r, c, "X,M,A,S") {
			part1 += 1
		}
		if G.MatchSym(r, c, "X   , M  ,  A ,   S") {
			part1 += 1
		}
		if G.MatchSym(r, c, "   X,  M , A  ,S   ") {
			part1 += 1
		}
	})
	timer.Part1(part1)

	part2 := 0
	G.ForEach(func(r int, c int, _ rune) {
		if G.MatchSym(r, c, "M  , A ,  S") && G.MatchSym(r, c, "  M, A ,S  ") {
			part2 += 1
		}
	})
	timer.Part2(part2)

	return part1, part2
}
