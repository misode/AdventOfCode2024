package day05

import (
	"fmt"
	"strings"

	"misode.dev/aoc-2024/utils"
)

func Solve() (int, int) {
	fmt.Println("=== Day 05 ===")

	lines := utils.ReadInput("in.txt")

	order := make(map[int][]int)
	updates := make([][]int, 0)
	first := true

	for _, line := range lines {
		if first {
			if len(line) == 0 {
				first = false
				continue
			}
			parts := strings.Split(line, "|")
			order[utils.StrToInt(parts[0])] = append(order[utils.StrToInt(parts[0])], utils.StrToInt(parts[1]))
		} else {
			parts := strings.Split(line, ",")
			update := make([]int, len(parts))
			for i, part := range parts {
				update[i] = utils.StrToInt(part)
			}
			updates = append(updates, update)
		}
	}

	part1 := 0
	part2 := 0
	for _, update := range updates {
		seq := utils.TopoSort(order, update)
		if utils.SliceEqual(update, seq) {
			part1 += seq[len(seq)/2]
		} else {
			part2 += seq[len(seq)/2]
		}
	}
	fmt.Println(part1)
	fmt.Println(part2)

	return part1, part2
}
