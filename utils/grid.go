package utils

import (
	"fmt"
	"strings"
)

type Grid[T comparable] [][]T

func MakeGrid(lines []string) Grid[rune] {
	grid := make([][]rune, len(lines))
	for i, line := range lines {
		grid[i] = []rune(line)
	}
	return grid
}

func MakeIntGrid(lines []string) Grid[int] {
	grid := make([][]int, len(lines))
	for i, line := range lines {
		row := make([]int, len(line))
		for j, char := range line {
			if char >= '0' && char <= '9' {
				row[j] = int(char - '0')
			}
		}
		grid[i] = row
	}
	return grid
}

func (grid *Grid[T]) Height() int {
	return len(*grid)
}

func (grid *Grid[T]) Width() int {
	return len((*grid)[0])
}

func (grid *Grid[T]) Transpose() Grid[T] {
	newGrid := make([][]T, grid.Width())
	for i := range newGrid {
		newGrid[i] = make([]T, grid.Height())
	}
	for r, row := range *grid {
		for c, val := range row {
			newGrid[c][r] = val
		}
	}
	return newGrid
}

func (grid *Grid[T]) ToLines() []string {
	lines := make([]string, grid.Height())
	for i, row := range *grid {
		builder := strings.Builder{}
		for _, val := range row {
			switch v := any(val).(type) {
			case rune:
				builder.WriteRune(v)
			default:
				builder.WriteString(fmt.Sprintf("%v", val))
			}
		}
		lines[i] = builder.String()
	}
	return lines
}

func (grid *Grid[T]) Print() {
	lines := grid.ToLines()
	for _, line := range lines {
		fmt.Println(line)
	}
}

func (grid *Grid[T]) IsInside(r int, c int) bool {
	return r >= 0 && c >= 0 && r < grid.Height() && c < grid.Width()
}

func (grid *Grid[T]) Is(r int, c int, check T) bool {
	return grid.IsInside(r, c) && (*grid)[r][c] == check
}

func (grid *Grid[T]) Mark(r int, c int, set T) bool {
	if grid.IsInside(r, c) {
		(*grid)[r][c] = set
		return true
	}
	return false
}

func (grid *Grid[T]) Get(r int, c int) (T, bool) {
	if grid.IsInside(r, c) {
		return (*grid)[r][c], true
	}
	var zero T
	return zero, false
}

func (grid *Grid[T]) Find(check T) (int, int, bool) {
	for r, row := range *grid {
		for c, val := range row {
			if val == check {
				return r, c, true
			}
		}
	}
	return 0, 0, false
}

func (grid *Grid[T]) ForEach(fn func(r int, c int, val T)) {
	for r, row := range *grid {
		for c, val := range row {
			fn(r, c, val)
		}
	}
}
