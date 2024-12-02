package day02

import (
	"fmt"
	"strconv"
	"strings"

	"misode.dev/aoc-2024/utils"
)

func Solve() {
	fmt.Println("=== Day 02 ===")

	lines := utils.ReadInput("day02/in.txt")

	part1 := 0
	part2 := 0

	for _, line := range lines {
		parts := strings.Split(line, " ")
		levels := make([]int, len(parts))
		for i, part := range parts {
			levels[i], _ = strconv.Atoi(part)
		}
		if CheckReport(levels) {
			part1 += 1
			part2 += 1
		} else {
			for i := range levels {
				var variation []int
				variation = append(variation, levels[:i]...)
				variation = append(variation, levels[i+1:]...)
				if CheckReport(variation) {
					part2 += 1
					break
				}
			}
		}
	}

	fmt.Println(part1)
	fmt.Println(part2)
}

func CheckReport(levels []int) bool {
	prev := -1
	dir := 0
	for _, lvl := range levels {
		if dir != 0 {
			if (lvl-prev < 0) != (dir < 0) {
				return false
			}
		}
		if prev > -1 {
			if prev == lvl || prev-lvl > 3 || prev-lvl < -3 {
				return false
			}
			dir = lvl - prev
		}
		prev = lvl
	}
	return true
}
