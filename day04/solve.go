package day04

import (
	"fmt"

	"misode.dev/aoc-2024/utils"
)

func Solve() (int, int) {
	fmt.Println("=== Day 04 ===")

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
	fmt.Println(part1)

	part2 := 0
	G.ForEach(func(r int, c int, _ rune) {
		if G.MatchSym(r, c, "M  , A ,  S") && G.MatchSym(r, c, "  M, A ,S  ") {
			part2 += 1
		}
	})
	fmt.Println(part2)

	return part1, part2
}
