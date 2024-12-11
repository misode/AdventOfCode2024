package utils

import (
	"fmt"
	"strings"
)

type Grid [][]rune

func MakeGrid(lines []string) Grid {
	grid := make([][]rune, len(lines))
	for i, line := range lines {
		grid[i] = []rune(line)
	}
	return grid
}

func (grid *Grid) Height() int {
	return len(*grid)
}

func (grid *Grid) Width() int {
	return len((*grid)[0])
}

func (grid *Grid) Transpose() Grid {
	newGrid := make([][]rune, grid.Width())
	for i := range newGrid {
		newGrid[i] = make([]rune, grid.Height())
	}
	for r, row := range *grid {
		for c, val := range row {
			newGrid[c][r] = val
		}
	}
	return newGrid
}

func (grid *Grid) ToLines() []string {
	lines := make([]string, grid.Height())
	for i, row := range *grid {
		lines[i] = string(row)
	}
	return lines
}

func (grid *Grid) Print() {
	lines := grid.ToLines()
	for _, line := range lines {
		fmt.Println(line)
	}
}

func (grid *Grid) IsInside(r int, c int) bool {
	return r >= 0 && c >= 0 && r < grid.Height() && c < grid.Width()
}

func (grid *Grid) Is(r int, c int, check rune) bool {
	return grid.IsInside(r, c) && (*grid)[r][c] == check
}

func (grid *Grid) Mark(r int, c int, set rune) bool {
	if grid.IsInside(r, c) {
		(*grid)[r][c] = set
		return true
	}
	return false
}

func (grid *Grid) Get(r int, c int) (rune, bool) {
	if grid.IsInside(r, c) {
		return (*grid)[r][c], true
	}
	return rune(0), false
}

func (grid *Grid) Match(r int, c int, search string) bool {
	subgrid := MakeGrid(strings.Split(search, ","))

	for dr := 0; dr < subgrid.Height(); dr++ {
		for dc := 0; dc < subgrid.Width(); dc++ {
			if !subgrid.Is(dr, dc, ' ') && !grid.Is(r+dr, c+dc, subgrid[dr][dc]) {
				return false
			}
		}
	}
	return true
}

func (grid *Grid) MatchSym(r int, c int, search string) bool {
	return grid.Match(r, c, search) || grid.Match(r, c, ReverseStr(search))
}

func (grid *Grid) Find(check rune) (int, int, bool) {
	for r, row := range *grid {
		for c, val := range row {
			if val == check {
				return r, c, true
			}
		}
	}
	return 0, 0, false
}

func (grid *Grid) ForEach(fn func(r int, c int, val rune)) {
	for r, row := range *grid {
		for c, val := range row {
			fn(r, c, val)
		}
	}
}
