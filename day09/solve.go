package day09

import (
	"strings"

	"misode.dev/aoc-2024/utils"
)

func Solve() (int, int) {
	timer := utils.StartTimer()

	lines := utils.ReadInput("in.txt")

	orig := make([]int, 0)
	for i, c := range strings.Split(lines[0], "") {
		for range utils.StrToInt(c) {
			if i%2 == 0 {
				orig = append(orig, i/2)
			} else {
				orig = append(orig, -1)
			}
		}
	}
	disk := make([]int, len(orig))
	timer.Parsed()

	copy(disk, orig)
	firstFree := 0
	for disk[firstFree] != -1 {
		firstFree += 1
	}
	for i := len(disk) - 1; i >= firstFree; i-- {
		disk[firstFree] = disk[i]
		disk[i] = -1
		for disk[firstFree] != -1 {
			firstFree += 1
		}
	}

	part1 := ComputeChecksum(disk)
	timer.Part(part1)

	copy(disk, orig)
	for i := len(disk) - 1; i > 0; i-- {
		n := disk[i]
		if n == -1 {
			continue
		}
		src := i
		for src-1 > 0 && disk[src-1] == n {
			src -= 1
		}
		chunkSize := i - src + 1
		foundsize := 0
		for dst := 0; dst < src; dst++ {
			if disk[dst] == -1 {
				foundsize += 1
				if foundsize == chunkSize {
					for j := range chunkSize {
						disk[dst-chunkSize+1+j] = n
						disk[src+j] = -1
					}
					break
				}
			} else {
				foundsize = 0
			}
		}
		i = src
	}

	part2 := ComputeChecksum(disk)
	timer.Part(part2)

	return part1, part2
}

func ComputeChecksum(disk []int) int {
	result := 0
	for i, n := range disk {
		if n != -1 {
			result += i * n
		}
	}
	return result
}
