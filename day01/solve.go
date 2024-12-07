package day01

import (
	"fmt"
	"sort"

	"misode.dev/aoc-2024/utils"
)

func Solve() (int, int) {
	fmt.Println("=== Day 01 ===")

	lines := utils.ReadInput("in.txt")

	left := make([]int, len(lines))
	right := make([]int, len(lines))
	for i, line := range lines {
		parts := utils.SplitInts(line, " ")
		left[i] = parts[0]
		right[i] = parts[1]
	}

	sort.IntSlice(left).Sort()
	sort.IntSlice(right).Sort()

	part1 := 0
	for i := range lines {
		l := left[i]
		r := right[i]
		diff := r - l
		if diff < 0 {
			diff = -diff
		}
		part1 += diff
	}
	fmt.Println(part1)

	counts := utils.Counter(right)

	part2 := 0
	for _, num := range left {
		part2 += num * counts[num]
	}
	fmt.Println(part2)

	return part1, part2
}
