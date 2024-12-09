package day05

import (
	"misode.dev/aoc-2024/utils"
)

func Solve() (int, int) {
	timer := utils.StartDay(5)

	lines := utils.ReadInput("in.txt")
	groups := utils.SplitLinesOnEmpty(lines)

	graph := make(map[int][]int)
	for _, line := range groups[0] {
		nums := utils.SplitInts(line, "|")
		graph[nums[0]] = append(graph[nums[0]], nums[1])
	}

	updates := make([][]int, 0)
	for _, line := range groups[1] {
		update := utils.SplitInts(line, ",")
		updates = append(updates, update)
	}
	timer.Parsed()

	part1 := 0
	part2 := 0
	for _, update := range updates {
		seq := utils.TopoSort(graph, update)
		if utils.SliceEqual(update, seq) {
			part1 += seq[len(seq)/2]
		} else {
			part2 += seq[len(seq)/2]
		}
	}
	timer.Parts(part1, part2)

	return part1, part2
}
