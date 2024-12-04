package day04

import (
	"fmt"

	"misode.dev/aoc-2024/utils"
)

func Solve() (int, int) {
	fmt.Println("=== Day 04 ===")

	lines := utils.ReadInput("in.txt")
	G := utils.MakeGrid(lines)
	N := len(G)
	utils.AssertEqual(N, len(G[0]))

	part1 := 0
	for r := 0; r < N; r++ {
		for c := 0; c < N; c++ {
			// Horizontal
			if utils.MatchSubGrid(G, r, c, "XMAS") {
				part1 += 1
			}
			if utils.MatchSubGrid(G, r, c, "SAMX") {
				part1 += 1
			}
			// Vertical
			if utils.MatchSubGrid(G, r, c, "X,M,A,S") {
				part1 += 1
			}
			if utils.MatchSubGrid(G, r, c, "S,A,M,X") {
				part1 += 1
			}
			// Primary diagonal
			if utils.MatchSubGrid(G, r, c, "X   , M  ,  A ,   S") {
				part1 += 1
			}
			if utils.MatchSubGrid(G, r, c, "S   , A  ,  M ,   X") {
				part1 += 1
			}
			// Secondary diagonal
			if utils.MatchSubGrid(G, r, c, "   X,  M , A  ,S   ") {
				part1 += 1
			}
			if utils.MatchSubGrid(G, r, c, "   S,  A , M  ,X   ") {
				part1 += 1
			}
		}
	}
	fmt.Println(part1)

	part2 := 0
	for r := 0; r+2 < N; r++ {
		for c := 0; c+2 < N; c++ {
			primary := (G[r][c] == 'M' && G[r+2][c+2] == 'S') || (G[r+2][c+2] == 'M' && G[r][c] == 'S')
			secondary := (G[r][c+2] == 'M' && G[r+2][c] == 'S') || (G[r+2][c] == 'M' && G[r][c+2] == 'S')
			if G[r+1][c+1] == 'A' && primary && secondary {
				part2 += 1
			}
		}
	}
	fmt.Println(part2)

	return part1, part2
}
