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
	G.ForEach(func(r int, c int) {
		if G.MatchSub(r, c, "XMAS") || G.MatchSub(r, c, "SAMX") {
			part1 += 1
		}
		if G.MatchSub(r, c, "X,M,A,S") || G.MatchSub(r, c, "S,A,M,X") {
			part1 += 1
		}
		if G.MatchSub(r, c, "X   , M  ,  A ,   S") || G.MatchSub(r, c, "S   , A  ,  M ,   X") {
			part1 += 1
		}
		if G.MatchSub(r, c, "   X,  M , A  ,S   ") || G.MatchSub(r, c, "   S,  A , M  ,X   ") {
			part1 += 1
		}
	})
	fmt.Println(part1)

	part2 := 0
	G.ForEach(func(r int, c int) {
		primary := (G.Is(r, c, 'M') && G.Is(r+2, c+2, 'S')) || (G.Is(r+2, c+2, 'M') && G.Is(r, c, 'S'))
		secondary := (G.Is(r, c+2, 'M') && G.Is(r+2, c, 'S')) || (G.Is(r+2, c, 'M') && G.Is(r, c+2, 'S'))
		if G.Is(r+1, c+1, 'A') && primary && secondary {
			part2 += 1
		}
	})
	fmt.Println(part2)

	return part1, part2
}
