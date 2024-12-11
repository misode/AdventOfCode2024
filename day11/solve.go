package day11

import (
	"strconv"

	"misode.dev/aoc-2024/utils"
)

func Solve() (int, int) {
	timer := utils.StartDay(11)

	lines := utils.ReadInput("in.txt")
	ints := utils.SplitInts(lines[0], " ")
	stones := utils.Counter(ints)

	part1 := simulate(stones, 25)
	timer.Part1(part1)
	part2 := simulate(stones, 75)
	timer.Part2(part2)

	return part1, part2
}

func simulate(stones map[int]int, blinks int) int {
	for range blinks {
		nextStones := make(map[int]int)
		for s, count := range stones {
			if s == 0 {
				nextStones[s+1] += count
				continue
			}
			str := strconv.Itoa(s)
			if len(str)%2 == 0 {
				nextStones[utils.StrToInt(str[0:len(str)/2])] += count
				nextStones[utils.StrToInt(str[len(str)/2:])] += count
			} else {
				nextStones[s*2024] += count
			}
		}
		stones = nextStones
	}
	totalCount := 0
	for _, count := range stones {
		totalCount += count
	}
	return totalCount
}
