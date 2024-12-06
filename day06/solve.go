package day06

import (
	"fmt"

	"misode.dev/aoc-2024/utils"
)

func Solve() (int, int) {

	fmt.Println("=== Day 06 ===")

	lines := utils.ReadInput("in.txt")
	grid := utils.MakeGrid(lines)

	guardR := 0
	guardC := 0
	dir := 0
	grid.ForEach(func(r int, c int) {
		if grid.Is(r, c, '^') {
			guardR = r
			guardC = c
		}
	})
	path := make(map[Point]bool)

	r, c := guardR, guardC

	path[Point{r: r, c: c}] = true

	for {
		blocked := dir == 0 && grid.Is(r-1, c, '#') ||
			dir == 1 && grid.Is(r, c+1, '#') ||
			dir == 2 && grid.Is(r+1, c, '#') ||
			dir == 3 && grid.Is(r, c-1, '#')
		if blocked {
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
		if grid.IsInside(r, c) {
			path[Point{r, c}] = true
		} else {
			break
		}
	}

	part1 := len(path)
	fmt.Println(part1)

	part2 := 0

	for point := range path {
		if point.r == guardR && point.c == guardC {
			continue
		}
		r, c := guardR, guardC
		dir := 0
		grid := utils.MakeGrid(lines)
		grid.Mark(point.r, point.c, '#')

		trail := make(map[PointDir]bool)
		trail[PointDir{r: r, c: c, dir: dir}] = true

		for {
			for {
				blocked := dir == 0 && grid.Is(r-1, c, '#') ||
					dir == 1 && grid.Is(r, c+1, '#') ||
					dir == 2 && grid.Is(r+1, c, '#') ||
					dir == 3 && grid.Is(r, c-1, '#')
				if blocked {
					dir = (dir + 1) % 4
				} else {
					break
				}
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
			cur := PointDir{r: r, c: c, dir: dir}
			_, ok := trail[cur]
			if ok {
				part2 += 1
				break
			} else if grid.IsInside(r, c) {
				trail[cur] = true
			} else {
				break
			}
		}
	}

	fmt.Println(part2)

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
