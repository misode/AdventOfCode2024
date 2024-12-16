package day16

import (
	"misode.dev/aoc-2024/utils"
)

var DIRS = []Point{{-1, 0}, {0, 1}, {1, 0}, {0, -1}}

func Solve() (int, int) {
	timer := utils.StartDay(16)

	lines := utils.ReadInput("in.txt")
	grid := utils.MakeGrid(lines)
	sr, sc, _ := grid.Find('S')
	er, ec, _ := grid.Find('E')

	visited := make(map[PointDir]bool)
	prevs := make(map[Node]map[Node]bool)
	pq := utils.MakeHeap(func(node *Node) int {
		return node.points
	})
	pq.Push(&Node{0, PointDir{sr, sc, 1}})

	part1 := -1
	part2 := -1
	for pq.Len() > 0 {
		node := pq.Pop()
		s := node.state
		if s.r == er && s.c == ec {
			part1 = node.points
			bestPath := make(map[Point]bool)
			queue := []*Node{node}
			for len(queue) > 0 {
				n := queue[len(queue)-1]
				queue = queue[:len(queue)-1]
				bestPath[Point{n.state.r, n.state.c}] = true
				for parent := range prevs[*n] {
					queue = append(queue, &parent)
				}
			}
			part2 = len(bestPath)
			break
		}

		tryEnqueue := func(points int, state *PointDir) {
			new := Node{points, *state}
			if prevs[new] == nil {
				prevs[new] = make(map[Node]bool)
			}
			prevs[new][*node] = true
			if !visited[new.state] {
				visited[new.state] = true
				pq.Push(&new)
			}
		}

		d := DIRS[s.dir]
		forward := PointDir{s.r + d.r, s.c + d.c, s.dir}
		if !grid.Is(forward.r, forward.c, '#') {
			tryEnqueue(node.points+1, &forward)
		}

		for _, rotate := range []int{-1, 1} {
			turn := PointDir{s.r, s.c, utils.Mod(s.dir+rotate, 4)}
			tryEnqueue(node.points+1000, &turn)
		}
	}
	timer.Parts(part1, part2)

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

type Node struct {
	points int
	state  PointDir
}
