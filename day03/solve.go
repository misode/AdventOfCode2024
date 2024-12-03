package day03

import (
	"fmt"
	"regexp"
	"strconv"

	"misode.dev/aoc-2024/utils"
)

func Solve() (int, int) {
	fmt.Println("=== Day 03 ===")

	lines := utils.ReadInput("in.txt")

	regex, err := regexp.Compile(`mul\((\d{1,3}),(\d{1,3})\)`)
	if err != nil {
		panic(err)
	}

	part1 := 0

	// regex.FindAllStringSubmatch(instr, -1)

	for _, line := range lines {
		matches := regex.FindAllStringSubmatch(line, -1)
		for _, match := range matches {
			// fmt.Println(match[0])
			num1, err1 := strconv.Atoi(match[1])
			num2, err2 := strconv.Atoi(match[2])
			if err1 != nil || err2 != nil {
				panic(err1)
			}
			part1 += num1 * num2
			// fmt.Println(match)
		}
	}
	fmt.Println(part1)

	regex2, err := regexp.Compile(`(?:mul\((\d{1,3}),(\d{1,3})\)|do\(\)|don't\(\))`)
	if err != nil {
		panic(err)
	}

	// lines = []string{"xmul(2,4)&mul[3,7]!^don't()_mul(5,5)+mul(32,64](mul(11,8)undo()?mul(8,5))"}

	part2 := 0
	enabled := true
	for _, line := range lines {
		matches := regex2.FindAllStringSubmatch(line, -1)
		for _, match := range matches {
			// fmt.Println(match)

			if match[0] == "do()" {
				enabled = true
			} else if match[0] == "don't()" {
				enabled = false
			} else if enabled {
				// fmt.Println(match[0])
				num1, err1 := strconv.Atoi(match[1])
				num2, err2 := strconv.Atoi(match[2])
				if err1 != nil || err2 != nil {
					panic(err1)
				}
				part2 += num1 * num2
			}
			// fmt.Println(match)
		}
	}

	fmt.Println(part2)

	return part1, part2
}
