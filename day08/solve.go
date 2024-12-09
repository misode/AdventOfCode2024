package day08

import (
	"fmt"

	"misode.dev/aoc-2024/utils"
)

func Solve() (int, int) {
	fmt.Println("=== Day 08 ===")

	lines := utils.ReadInput("in.txt")
	grid := utils.MakeGrid(lines)

	antennas := make(map[rune][]Point)
	grid.ForEach(func(r, c int, val rune) {
		if val != '.' {
			antennas[val] = append(antennas[val], Point{r, c})
		}
	})

	antinodes1 := make(map[Point]bool)
	antinodes2 := make(map[Point]bool)
	for _, nodes := range antennas {
		utils.ForCombinations(nodes, func(p1 Point, p2 Point) {
			for _, move := range []Move{MoveFromTo(p1, p2), MoveFromTo(p2, p1)} {
				move = move.Step()
				move = move.Step()
				if grid.IsInside(move.pos.r, move.pos.c) {
					antinodes1[move.pos] = true
				}
				for grid.IsInside(move.pos.r, move.pos.c) {
					antinodes2[move.pos] = true
					move = move.Step()
				}
			}
			antinodes2[p1] = true
			antinodes2[p2] = true
		})
	}
	part1 := len(antinodes1)
	fmt.Println(part1)

	part2 := len(antinodes2)
	fmt.Println(part2)

	return part1, part2
}

type Point struct {
	r int
	c int
}

type Move struct {
	pos Point
	dir Point
}

func MoveFromTo(from Point, to Point) Move {
	dir := Point{r: to.r - from.r, c: to.c - from.c}
	return Move{pos: from, dir: dir}
}

func (m *Move) To() Point {
	return Point{r: m.pos.r + m.dir.r, c: m.pos.c + m.dir.c}
}

func (m *Move) Step() Move {
	return Move{pos: m.To(), dir: m.dir}
}
