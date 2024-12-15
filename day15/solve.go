package day15

import (
	"slices"
	"strings"

	"misode.dev/aoc-2024/utils"
)

func Solve() (int, int) {
	timer := utils.StartDay(15)

	lines := utils.ReadInput("in.txt")
	groups := utils.SplitLinesOnEmpty(lines)
	grid := utils.MakeGrid(groups[0])
	path := strings.Join(groups[1], "")

	pr, pc, _ := grid.Find('@')

	dirs := map[rune]Point{'^': {-1, 0}, '>': {0, 1}, 'v': {1, 0}, '<': {0, -1}}
	for _, instr := range path {
		dir := dirs[instr]
		tr, tc := pr+dir.r, pc+dir.c
		for grid.Is(tr, tc, 'O') {
			tr += dir.r
			tc += dir.c
		}
		if !grid.Is(tr, tc, '#') {
			grid[tr][tc] = 'O'
			grid[pr][pc] = '.'
			grid[pr+dir.r][pc+dir.c] = '@'
			pr += dir.r
			pc += dir.c
		}
	}

	part1 := 0
	grid.ForEach(func(r, c int, val rune) {
		if val == 'O' {
			part1 += 100*r + c
		}
	})
	timer.Part1(part1)

	wider := make([]string, len(groups[0]))
	for i, line := range groups[0] {
		wide := strings.Builder{}
		for _, char := range line {
			if char == '#' {
				wide.WriteString("##")
			} else if char == 'O' {
				wide.WriteString("[]")
			} else if char == '.' {
				wide.WriteString("..")
			} else if char == '@' {
				wide.WriteString("@.")
			}
		}
		wider[i] = wide.String()
	}
	grid = utils.MakeGrid(wider)
	pr, pc, _ = grid.Find('@')

	for _, instr := range path {
		dir := dirs[instr]
		push := map[Point]bool{{pr, pc}: true}
		pushes := map[Point]bool{}
		stop := false
		for {
			nextPush := make(map[Point]bool)
			for p := range push {
				if grid.Is(p.r, p.c, '#') {
					stop = true
					break
				}
				tr, tc := p.r+dir.r, p.c+dir.c
				pushes[Point{tr, tc}] = true
				if grid.Is(tr, tc, '[') {
					if dir.c == 0 {
						pushes[Point{tr, tc + 1}] = true
						nextPush[Point{tr, tc}] = true
						nextPush[Point{tr, tc + 1}] = true
					} else {
						pushes[Point{tr, tc + dir.c}] = true
						nextPush[Point{tr, tc + dir.c}] = true
					}
				} else if grid.Is(tr, tc, ']') {
					if dir.c == 0 {
						pushes[Point{tr, tc - 1}] = true
						nextPush[Point{tr, tc}] = true
						nextPush[Point{tr, tc - 1}] = true
					} else {
						pushes[Point{tr, tc + dir.c}] = true
						nextPush[Point{tr, tc + dir.c}] = true
					}
				} else if grid.Is(tr, tc, '#') {
					nextPush[Point{tr, tc}] = true
					stop = true
				} else {
					delete(pushes, Point{tr, tc})
				}
			}
			push = nextPush
			if stop || len(push) == 0 {
				break
			}
		}
		if !stop && len(push) == 0 {
			ordered := make([]Point, 0)
			for p := range pushes {
				ordered = append(ordered, p)
			}
			slices.SortStableFunc(ordered, func(a, b Point) int {
				if dir.r == 0 {
					return (b.c - a.c) * dir.c
				} else {
					return (b.r - a.r) * dir.r
				}
			})
			for _, p := range ordered {
				grid[p.r+dir.r][p.c+dir.c] = grid[p.r][p.c]
				grid[p.r][p.c] = '.'
			}
			grid[pr+dir.r][pc+dir.c] = '@'
			grid[pr][pc] = '.'
			pr += dir.r
			pc += dir.c
		}
	}

	part2 := 0
	grid.ForEach(func(r, c int, val rune) {
		if val == '[' {
			part2 += 100*r + c
		}
	})
	timer.Part2(part2)

	return part1, part2
}

type Point struct {
	r int
	c int
}
