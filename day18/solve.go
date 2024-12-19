package day18

import (
	"strconv"

	"misode.dev/aoc-2024/utils"
)

func Solve() (int, string) {
	timer := utils.StartTimer()

	lines := utils.ReadInput("in.txt")
	points := make([]Point, len(lines))
	for i, line := range lines {
		x, y := utils.SplitInts2(line, ",")
		points[i] = Point{x, y}
	}
	timer.Parsed()
	S := 71

	grid := make([][]bool, S)
	for i := range S {
		grid[i] = make([]bool, S)
	}
	for i, p := range points {
		if i < 1024 {
			grid[p.x][p.y] = true
		}
	}
	part1, _ := solve(&grid)
	timer.Part(part1)

	part2 := ""
	for i := range S {
		grid[i] = make([]bool, S)
	}
	for _, p := range points {
		grid[p.x][p.y] = true
		_, possible := solve(&grid)
		if !possible {
			part2 = strconv.Itoa(p.x) + "," + strconv.Itoa(p.y)
			break
		}
	}
	timer.Part(part2)

	return part1, part2
}

var DIRS = [4]Point{{-1, 0}, {0, 1}, {1, 0}, {0, -1}}

type Point struct {
	x int
	y int
}

func solve(grid *[][]bool) (int, bool) {
	S := len(*grid)
	steps := -1
	solved := false
	visited := make(map[Point]bool)
	queue := []Point{{0, 0}}
	for !solved && len(queue) > 0 {
		steps += 1
		nextQueue := []Point{}
		for _, n := range queue {
			if n.x == S-1 && n.y == S-1 {
				solved = true
				break
			}
			for _, d := range DIRS {
				tx, ty := n.x+d.x, n.y+d.y
				if tx >= 0 && ty >= 0 && tx < S && ty < S {
					if !(*grid)[tx][ty] && !visited[Point{tx, ty}] {
						visited[Point{tx, ty}] = true
						nextQueue = append(nextQueue, Point{tx, ty})
					}
				}
			}
		}
		queue = nextQueue
	}
	return steps, solved
}
