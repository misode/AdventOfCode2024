package day13

import (
	"strings"

	"misode.dev/aoc-2024/utils"
)

func Solve() (int, int) {
	timer := utils.StartTimer()

	lines := utils.ReadInput("in.txt")
	groups := utils.SplitLinesOnEmpty(lines)

	part1 := 0
	for _, group := range groups {
		ax, ay, bx, by, px, py := ParseMachine(group)
		for n := 0; n <= 100; n++ {
			r := (px - (n * ax))
			if r%bx == 0 {
				m := r / bx
				if n*ay+m*by == py {
					part1 += 3*n + 1*m
					break
				}
			}
		}
	}
	timer.Part(part1)

	part2 := 0
	for _, group := range groups {
		ax, ay, bx, by, px, py := ParseMachine(group)

		px += 10000000000000
		py += 10000000000000

		// Find x in A*x = P  ->  x = A-1*P
		det := ax*by - ay*bx
		if det == 0 {
			continue
		}

		// n = (by/det)*px + (-bx/det)*py
		// m = (-ay/det)*px + (ax/det)*py
		// ->
		// n = (by*px - bx*py) / det
		// m = (-ay*px + ax*py) / det
		n := by*px - bx*py
		m := -ay*px + ax*py
		if n%det == 0 && m%det == 0 {
			n /= det
			m /= det
			part2 += 3*n + m
		}
	}
	timer.Part(part2)

	return part1, part2
}

func ParseMachine(group []string) (int, int, int, int, int, int) {
	aParts := strings.Split(group[0][len("Button A: "):], ", ")
	ax, ay := utils.StrToInt(aParts[0][2:]), utils.StrToInt(aParts[1][2:])
	bParts := strings.Split(group[1][len("Button B: "):], ", ")
	bx, by := utils.StrToInt(bParts[0][2:]), utils.StrToInt(bParts[1][2:])
	pParts := strings.Split(group[2][len("Prize: "):], ", ")
	px, py := utils.StrToInt(pParts[0][2:]), utils.StrToInt(pParts[1][2:])
	return ax, ay, bx, by, px, py
}
