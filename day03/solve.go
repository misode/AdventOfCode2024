package day03

import (
	"fmt"
	"strings"

	"misode.dev/aoc-2024/utils"
)

func Solve() (int, int) {
	fmt.Println("=== Day 03 ===")

	lines := utils.ReadInput("in.txt")
	memory := strings.Join(lines, "")

	part1 := 0
	for _, match := range utils.FindMatches(`mul\((\d{1,3}),(\d{1,3})\)`, memory) {
		part1 += utils.StrToInt(match[1]) * utils.StrToInt(match[2])
	}
	fmt.Println(part1)

	part2 := 0
	enabled := true
	for _, match := range utils.FindMatches(`mul\((\d{1,3}),(\d{1,3})\)|do\(\)|don't\(\)`, memory) {
		if match[0] == "do()" {
			enabled = true
		} else if match[0] == "don't()" {
			enabled = false
		} else if enabled {
			part2 += utils.StrToInt(match[1]) * utils.StrToInt(match[2])
		}
	}
	fmt.Println(part2)

	return part1, part2
}
