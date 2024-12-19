package day15

import (
	"maps"
	"slices"
	"strings"

	"misode.dev/aoc-2024/utils"
)

var DIRS = map[rune]Point{'^': {-1, 0}, '>': {0, 1}, 'v': {1, 0}, '<': {0, -1}}

func Solve() (int, int) {
	timer := utils.StartTimer()

	lines := utils.ReadInput("in.txt")
	groups := utils.SplitLinesOnEmpty(lines)
	grid := utils.MakeGrid(groups[0])
	path := strings.Join(groups[1], "")
	pr, pc, _ := grid.Find('@')

	for _, instr := range path {
		dir := DIRS[instr]
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
	timer.Part(part1)

	wider := make([]string, len(groups[0]))
	for i, line := range groups[0] {
		line = strings.ReplaceAll(line, "#", "##")
		line = strings.ReplaceAll(line, "O", "[]")
		line = strings.ReplaceAll(line, ".", "..")
		line = strings.ReplaceAll(line, "@", "@.")
		wider[i] = line
	}
	grid = utils.MakeGrid(wider)
	pr, pc, _ = grid.Find('@')

	for _, instr := range path {
		dir := DIRS[instr]
		push := map[Point]bool{{pr, pc}: true}
		moves := map[Point]bool{}
		wall := false
		for !wall && len(push) > 0 {
			nextPush := make(map[Point]bool)
			for p := range push {
				if grid.Is(p.r, p.c, '#') {
					wall = true
					break
				}
				tr, tc := p.r+dir.r, p.c+dir.c
				if grid.Is(tr, tc, '.') {
					continue
				}
				moves[Point{tr, tc}] = true
				if dir.r != 0 { // Up or down
					nextPush[Point{tr, tc}] = true
					if grid.Is(tr, tc, '[') {
						nextPush[Point{tr, tc + 1}] = true
						moves[Point{tr, tc + 1}] = true
					} else {
						nextPush[Point{tr, tc - 1}] = true
						moves[Point{tr, tc - 1}] = true
					}
				} else { // Left or right
					moves[Point{tr, tc + dir.c}] = true
					nextPush[Point{tr, tc + dir.c}] = true
				}
			}
			push = nextPush
		}
		if !wall {
			orderedMoves := slices.Collect(maps.Keys(moves))
			slices.SortFunc(orderedMoves, func(a, b Point) int {
				return (b.c-a.c)*dir.c + (b.r-a.r)*dir.r
			})
			for _, p := range orderedMoves {
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
	timer.Part(part2)

	return part1, part2
}

type Point struct {
	r int
	c int
}
