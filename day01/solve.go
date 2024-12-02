package day01

import (
	"fmt"
	"sort"
	"strconv"
	"strings"

	"misode.dev/aoc-2024/utils"
)

func Solve() {
	fmt.Println("=== Day 01 ===")

	lines := utils.ReadInput("day01/in.txt")

	left := make([]int, len(lines))
	right := make([]int, len(lines))
	for i, line := range lines {
		parts := strings.Split(line, "   ")
		left[i], _ = strconv.Atoi(parts[0])
		right[i], _ = strconv.Atoi(parts[1])
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

	counts := make(map[int]int)
	for _, num := range right {
		counts[num] = counts[num] + 1
	}

	part2 := 0
	for _, num := range left {
		count := counts[num]
		part2 += num * count
	}
	fmt.Println(part2)
}
