package day10

import (
	"misode.dev/aoc-2024/utils"
)

func Solve() (int, int) {
	timer := utils.StartTimer()

	lines := utils.ReadInput("in.txt")
	grid := utils.MakeIntGrid(lines)

	part1 := 0
	part2 := 0
	grid.ForEach(func(r int, c int, val int) {
		if val != 0 {
			return
		}
		trails := make(map[Point]int)
		FindTrails(&grid, r, c, val, trails)
		part1 += len(trails)
		for _, score := range trails {
			part2 += score
		}
	})
	timer.Parts(part1, part2)

	return part1, part2
}

var DIRS = [4]Point{{-1, 0}, {0, 1}, {1, 0}, {0, -1}}

func FindTrails(grid *utils.Grid[int], r int, c int, val int, trails map[Point]int) {
	if val == 9 {
		trails[Point{r, c}] += 1
		return
	}
	for _, dir := range DIRS {
		rr, cc := r+dir.r, c+dir.c
		if grid.Is(rr, cc, val+1) {
			FindTrails(grid, rr, cc, val+1, trails)
		}
	}
}

type Point struct {
	r int
	c int
}
