package day14

import (
	"strings"

	"misode.dev/aoc-2024/utils"
)

func Solve() (int, int) {
	timer := utils.StartDay(14)

	lines := utils.ReadInput("in.txt")

	robots := make([]Robot, len(lines))
	for i, line := range lines {
		parts := strings.Split(line, " ")
		px, py := utils.SplitInts2(parts[0][2:], ",")
		vx, vy := utils.SplitInts2(parts[1][2:], ",")
		robots[i] = Robot{px, py, vx, vy}
	}

	W, H, S := 101, 103, 100
	q1, q2, q3, q4 := 0, 0, 0, 0
	for _, r := range robots {
		PX := utils.Mod(r.px+r.vx*S, W)
		PY := utils.Mod(r.py+r.vy*S, H)
		if PX < W/2 && PY < H/2 {
			q1 += 1
		} else if PX > W/2 && PY < H/2 {
			q2 += 1
		} else if PX < W/2 && PY > H/2 {
			q3 += 1
		} else if PX > W/2 && PY > H/2 {
			q4 += 1
		}
	}
	part1 := q1 * q2 * q3 * q4
	timer.Part1(part1)

	part2 := 0
	for {
		part2 += 1
		ok := true
		uniqueCheck := make(map[Point]bool)
		for i, r := range robots {
			px := utils.Mod(r.px+r.vx, W)
			py := utils.Mod(r.py+r.vy, H)
			robots[i].px = px
			robots[i].py = py
			p := Point{px, py}
			if ok && uniqueCheck[p] {
				ok = false
			}
			uniqueCheck[p] = true
		}
		if ok {
			break
		}
	}
	timer.Part2(part2)

	return part1, part2
}

type Robot struct {
	px int
	py int
	vx int
	vy int
}

type Point struct {
	x int
	y int
}
