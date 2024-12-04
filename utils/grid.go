package utils

import "strings"

type Grid struct {
	data [][]rune
}

func MakeGrid(lines []string) Grid {
	data := make([][]rune, len(lines))
	for i, line := range lines {
		data[i] = []rune(line)
	}
	return Grid{data}
}

func (g *Grid) Height() int {
	return len(g.data)
}

func (g *Grid) Width() int {
	return len(g.data[0])
}

func (g *Grid) Transpose() Grid {
	data := make([][]rune, g.Width())
	for i := range data {
		data[i] = make([]rune, g.Height())
	}
	for r, row := range g.data {
		for c, val := range row {
			data[c][r] = val
		}
	}
	return Grid{data}
}

func (g *Grid) ToLines() []string {
	lines := make([]string, g.Height())
	for i, row := range g.data {
		lines[i] = string(row)
	}
	return lines
}

func (g *Grid) Is(r int, c int, check rune) bool {
	return r < g.Height() && c < g.Width() && g.data[r][c] == check
}

func (g *Grid) Match(r int, c int, search string) bool {
	subgrid := MakeGrid(strings.Split(search, ","))

	for dr := 0; dr < subgrid.Height(); dr++ {
		for dc := 0; dc < subgrid.Width(); dc++ {
			if !subgrid.Is(dr, dc, ' ') && !g.Is(r+dr, c+dc, subgrid.data[dr][dc]) {
				return false
			}
		}
	}
	return true
}

func (g *Grid) MatchSym(r int, c int, search string) bool {
	return g.Match(r, c, search) || g.Match(r, c, ReverseStr(search))
}

func (g *Grid) ForEach(fn func(r int, c int)) {
	for r := 0; r < g.Height(); r++ {
		for c := 0; c < g.Width(); c++ {
			fn(r, c)
		}
	}
}
