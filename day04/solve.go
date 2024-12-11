package day04

import (
	"strings"

	"misode.dev/aoc-2024/utils"
)

func Solve() (int, int) {
	timer := utils.StartDay(4)

	lines := utils.ReadInput("in.txt")
	grid := utils.MakeGrid(lines)

	part1 := 0
	grid.ForEach(func(r int, c int, _ rune) {
		if matchSym(&grid, r, c, "XMAS") {
			part1 += 1
		}
		if matchSym(&grid, r, c, "X,M,A,S") {
			part1 += 1
		}
		if matchSym(&grid, r, c, "X   , M  ,  A ,   S") {
			part1 += 1
		}
		if matchSym(&grid, r, c, "   X,  M , A  ,S   ") {
			part1 += 1
		}
	})
	timer.Part1(part1)

	part2 := 0
	grid.ForEach(func(r int, c int, _ rune) {
		if matchSym(&grid, r, c, "M  , A ,  S") && matchSym(&grid, r, c, "  M, A ,S  ") {
			part2 += 1
		}
	})
	timer.Part2(part2)

	return part1, part2
}

func match(grid *utils.Grid[rune], r int, c int, search string) bool {
	subgrid := utils.MakeGrid(strings.Split(search, ","))

	for dr := 0; dr < subgrid.Height(); dr++ {
		for dc := 0; dc < subgrid.Width(); dc++ {
			if !subgrid.Is(dr, dc, ' ') && !grid.Is(r+dr, c+dc, subgrid[dr][dc]) {
				return false
			}
		}
	}
	return true
}

func matchSym(grid *utils.Grid[rune], r int, c int, search string) bool {
	return match(grid, r, c, search) || match(grid, r, c, utils.ReverseStr(search))
}
