package day10

import (
	"misode.dev/aoc-2024/utils"
)

func Solve() (int, int) {
	timer := utils.StartDay(10)

	lines := utils.ReadInput("in.txt")
	grid := make([][]int, len(lines))
	for i, line := range lines {
		grid[i] = make([]int, 0)
		for _, char := range line {
			grid[i] = append(grid[i], utils.StrToInt(string(char)))
		}
	}

	part1 := 0
	part2 := 0
	for r, row := range grid {
		for c, val := range row {
			if val != 0 {
				continue
			}
			trails := make(map[Point]int)
			FindTrails(grid, r, c, val, trails)
			part1 += len(trails)
			for _, score := range trails {
				part2 += score
			}
		}
	}
	timer.Parts(part1, part2)

	return part1, part2
}

var DIRS = [4]Point{{-1, 0}, {0, 1}, {1, 0}, {0, -1}}

func FindTrails(grid [][]int, r int, c int, val int, trails map[Point]int) {
	if val == 9 {
		trails[Point{r, c}] += 1
		return
	}
	for _, dir := range DIRS {
		rr, cc := r+dir.r, c+dir.c
		if rr >= 0 && cc >= 0 && rr < len(grid) && cc < len(grid[0]) && grid[rr][cc] == val+1 {
			FindTrails(grid, rr, cc, val+1, trails)
		}
	}
}

type Point struct {
	r int
	c int
}
