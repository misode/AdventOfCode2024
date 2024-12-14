package day01

import (
	"sort"

	"misode.dev/aoc-2024/utils"
)

func Solve() (int, int) {
	timer := utils.StartDay(1)

	lines := utils.ReadInput("in.txt")

	left := make([]int, len(lines))
	right := make([]int, len(lines))
	for i, line := range lines {
		left[i], right[i] = utils.SplitInts2(line, " ")
	}
	timer.Parsed()

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
	timer.Part1(part1)

	counts := utils.Counter(right)

	part2 := 0
	for _, num := range left {
		part2 += num * counts[num]
	}
	timer.Part2(part2)

	return part1, part2
}
