package day06

import (
	"misode.dev/aoc-2024/utils"
)

func Solve() (int, int) {
	timer := utils.StartTimer()

	lines := utils.ReadInput("in.txt")
	grid := utils.MakeGrid(lines)

	guardR, guardC, _ := grid.Find('^')

	path, _ := simulate(&grid, guardR, guardC, 0)
	part1 := len(path)
	timer.Part(part1)

	part2 := 0
	for point := range path {
		if point.r == guardR && point.c == guardC {
			continue
		}
		grid := utils.MakeGrid(lines)
		grid[point.r][point.c] = '#'
		_, cycle := simulate(&grid, guardR, guardC, 0)
		if cycle {
			part2 += 1
		}
	}
	timer.Part(part2)

	return part1, part2
}

type Point struct {
	r int
	c int
}

type PointDir struct {
	r   int
	c   int
	dir int
}

func simulate(grid *utils.Grid[rune], r int, c int, dir int) (map[Point]bool, bool) {
	path := make(map[Point]bool)
	trail := make(map[PointDir]bool)

	for {
		path[Point{r, c}] = true
		trail[PointDir{r, c, dir}] = true
		for dir == 0 && grid.Is(r-1, c, '#') ||
			dir == 1 && grid.Is(r, c+1, '#') ||
			dir == 2 && grid.Is(r+1, c, '#') ||
			dir == 3 && grid.Is(r, c-1, '#') {
			dir = (dir + 1) % 4
		}
		switch dir {
		case 0:
			r -= 1
		case 1:
			c += 1
		case 2:
			r += 1
		case 3:
			c -= 1
		}
		_, ok := trail[PointDir{r, c, dir}]
		if ok {
			return path, true // Cycle
		} else if !grid.IsInside(r, c) {
			return path, false // Outside
		}
	}
}
