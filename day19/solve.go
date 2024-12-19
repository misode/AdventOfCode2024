package day19

import (
	"strings"

	"misode.dev/aoc-2024/utils"
)

func Solve() (int, int) {
	timer := utils.StartDay(19)

	lines := utils.ReadInput("in.txt")
	groups := utils.SplitLinesOnEmpty(lines)
	patterns := strings.Split(groups[0][0], ", ")
	designs := groups[1]
	timer.Parsed()

	cache := map[string]int{}

	var search func(design string) int
	search = func(design string) int {
		if len(design) == 0 {
			return 1
		}
		count, ok := cache[design]
		if ok {
			return count
		}
		for _, pattern := range patterns {
			rest, ok := strings.CutPrefix(design, pattern)
			if ok {
				count += search(rest)
			}
		}
		cache[design] = count
		return count
	}

	part1 := 0
	part2 := 0
	for _, design := range designs {
		count := search(design)
		if count > 0 {
			part1 += 1
		}
		part2 += count
	}
	timer.Parts(part1, part2)

	return part1, part2
}
