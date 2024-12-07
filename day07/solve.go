package day07

import (
	"fmt"
	"slices"
	"strconv"
	"strings"

	"misode.dev/aoc-2024/utils"
)

func Solve() (int, int) {
	fmt.Println("=== Day 07 ===")

	lines := utils.ReadInput("in.txt")

	part1 := 0
	part2 := 0
	for _, line := range lines {
		parts := strings.Split(line, ": ")
		result := utils.StrToInt(parts[0])
		nums := utils.SplitInts(parts[1], " ")
		if slices.Contains(Solutions(nums[0], nums[1:], false), result) {
			part1 += result
		}
		if slices.Contains(Solutions(nums[0], nums[1:], true), result) {
			part2 += result
		}
	}
	fmt.Println(part1)
	fmt.Println(part2)

	return int(part1), part2
}

func Solutions(cur int, nums []int, part2 bool) []int {
	if len(nums) == 0 {
		return []int{cur}
	}
	next, rest := nums[0], nums[1:]
	solutions := slices.Concat(
		Solutions(cur+next, rest, part2),
		Solutions(cur*next, rest, part2))
	if part2 {
		concat := utils.StrToInt(strconv.Itoa(cur) + strconv.Itoa(next))
		solutions = slices.Concat(solutions, Solutions(concat, rest, part2))
	}
	return solutions
}
