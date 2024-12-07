package day07

import (
	"fmt"
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
		if Solution(nums[0], nums[1:], result, false) {
			part1 += result
		}
		if Solution(nums[0], nums[1:], result, true) {
			part2 += result
		}
	}
	fmt.Println(part1)
	fmt.Println(part2)

	return int(part1), part2
}

func Solution(cur int, nums []int, result int, part2 bool) bool {
	if len(nums) == 0 {
		return cur == result
	}
	if cur > result {
		return false
	}
	next, rest := nums[0], nums[1:]
	if Solution(cur+next, rest, result, part2) {
		return true
	}
	if Solution(cur*next, rest, result, part2) {
		return true
	}
	if part2 {
		concat := utils.StrToInt(strconv.Itoa(cur) + strconv.Itoa(next))
		if Solution(concat, rest, result, part2) {
			return true
		}
	}
	return false
}