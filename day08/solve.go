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
			a1 := Point{r: p1.r - (p2.r - p1.r), c: p1.c - (p2.c - p1.c)}
			if grid.IsInside(a1.r, a1.c) {
				antinodes1[a1] = true
			}
			for grid.IsInside(a1.r, a1.c) {
				antinodes2[a1] = true
				a1.r -= (p2.r - p1.r)
				a1.c -= (p2.c - p1.c)
			}
			a2 := Point{r: p2.r - (p1.r - p2.r), c: p2.c - (p1.c - p2.c)}
			if grid.IsInside(a2.r, a2.c) {
				antinodes1[a2] = true
			}
			for grid.IsInside(a2.r, a2.c) {
				antinodes2[a2] = true
				a2.r -= (p1.r - p2.r)
				a2.c -= (p1.c - p2.c)
			}
			antinodes2[p1] = true
			antinodes2[p2] = true
		})
	}
	part1 := len(antinodes1)
	fmt.Println(part1)

	part2 := len(antinodes2)
	fmt.Println(part2)

	return int(part1), part2
}

type Point struct {
	r int
	c int
}
