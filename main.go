package main

import (
	"flag"
	"fmt"

	"misode.dev/aoc-2024/day01"
	"misode.dev/aoc-2024/day02"
	"misode.dev/aoc-2024/day03"
	"misode.dev/aoc-2024/day04"
	"misode.dev/aoc-2024/day05"
	"misode.dev/aoc-2024/day06"
	"misode.dev/aoc-2024/day07"
	"misode.dev/aoc-2024/day08"
	"misode.dev/aoc-2024/day09"
	"misode.dev/aoc-2024/day10"
	"misode.dev/aoc-2024/day11"
	"misode.dev/aoc-2024/day12"
	"misode.dev/aoc-2024/day13"
	"misode.dev/aoc-2024/day14"
	"misode.dev/aoc-2024/day15"
	"misode.dev/aoc-2024/day16"
	"misode.dev/aoc-2024/day17"
	"misode.dev/aoc-2024/day18"
	"misode.dev/aoc-2024/day19"
	"misode.dev/aoc-2024/utils"
)

var DAYS = []func(){
	func() { day01.Solve() },
	func() { day02.Solve() },
	func() { day03.Solve() },
	func() { day04.Solve() },
	func() { day05.Solve() },
	func() { day06.Solve() },
	func() { day07.Solve() },
	func() { day08.Solve() },
	func() { day09.Solve() },
	func() { day10.Solve() },
	func() { day11.Solve() },
	func() { day12.Solve() },
	func() { day13.Solve() },
	func() { day14.Solve() },
	func() { day15.Solve() },
	func() { day16.Solve() },
	func() { day17.Solve() },
	func() { day18.Solve() },
	func() { day19.Solve() },
}

func main() {
	day := flag.Int("day", 0, "The day to solve")
	flag.Parse()

	timer := utils.StartTimer()
	for i, solver := range DAYS {
		if *day == 0 || *day-1 == i {
			fmt.Printf("=== Day %02d ===\n", i+1)
			solver()
		}
	}
	fmt.Printf("=== Total %v ===\n", timer.Diff())
}
