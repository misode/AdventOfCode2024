package day12

import (
	"misode.dev/aoc-2024/utils"
)

func Solve() (int, int) {
	timer := utils.StartDay(12)

	lines := utils.ReadInput("in.txt")
	grid := utils.MakeGrid(lines)

	part1 := CalculatePrice(&grid, false)
	timer.Part1(part1)

	part2 := CalculatePrice(&grid, true)
	timer.Part2(part2)

	return part1, part2
}

func CalculatePrice(grid *utils.Grid[rune], hasDiscount bool) int {
	accounted := make(map[Point]bool)
	visited := make(map[Point]bool)
	perims := make(map[PointDir]bool)

	var flood func(r int, c int) (int, int)
	flood = func(r int, c int) (int, int) {
		if visited[Point{r, c}] {
			return 0, 0
		}
		visited[Point{r, c}] = true
		cur := (*grid)[r][c]
		area := 1
		perim := 0
		for i, d := range DIRS {
			rr, cc := r+d.r, c+d.c
			if grid.Is(rr, cc, cur) {
				a, p := flood(rr, cc)
				area += a
				perim += p
			} else if hasDiscount {
				sides := [2]Point{DIRS[(i+1)%4], DIRS[(i+3)%4]}
				n1 := PointDir{r + sides[0].r, c + sides[0].c, i}
				n2 := PointDir{r + sides[1].r, c + sides[1].c, i}
				if !perims[n1] && !perims[n2] {
					perim += 1 // Count new perimeter
				} else if perims[n1] && perims[n2] {
					perim -= 1 // Merge two perimeters
				}
				perims[PointDir{r, c, i}] = true
			} else {
				perim += 1
			}
		}
		accounted[Point{r, c}] = true
		return area, perim
	}

	price := 0
	grid.ForEach(func(r, c int, val rune) {
		if accounted[Point{r, c}] {
			return
		}
		area, perim := flood(r, c)
		price += area * perim
		clear(visited)
		clear(perims)
	})
	return price
}

var DIRS = [4]Point{{-1, 0}, {0, 1}, {1, 0}, {0, -1}}

type Point struct {
	r int
	c int
}

type PointDir struct {
	r   int
	c   int
	dir int
}
