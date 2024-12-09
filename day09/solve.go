package day09

import (
	"fmt"
	"strings"

	"misode.dev/aoc-2024/utils"
)

func Solve() (int, int) {
	fmt.Println("=== Day 09 ===")

	lines := utils.ReadInput("in.txt")
	line := lines[0]

	orig := make([]int, 0)
	for i, c := range strings.Split(line, "") {
		n := utils.StrToInt(c)
		for j := 0; j < n; j++ {
			if i%2 == 0 {
				orig = append(orig, i/2)
			} else {
				orig = append(orig, -1)
			}
		}
	}
	disk := make([]int, len(orig))
	copy(disk, orig)

	p := 0
	for disk[p] != -1 {
		p += 1
	}

	for i := len(disk) - 1; i >= p; i-- {
		disk[p] = disk[i]
		disk[i] = -1
		for disk[p] != -1 {
			p += 1
		}
	}

	part1 := 0
	for i, n := range disk {
		if n != -1 {
			part1 += i * n
		}
	}
	fmt.Println(part1)

	copy(disk, orig)
	for i := len(disk) - 1; i > 0; i-- {
		n := disk[i]
		if n == -1 {
			continue
		}
		j := i
		for j-1 > 0 && disk[j-1] == n {
			j -= 1
		}
		size := i - j + 1
		s := 0
		for k := 0; k < j; k++ {
			if disk[k] == -1 {
				s += 1
				if s >= size {
					for l := 0; l < size; l++ {
						disk[k-s+1+l] = n
						disk[j+l] = -1
					}
					break
				}
			} else {
				s = 0
			}
		}
		i = j
	}

	part2 := 0
	for i, n := range disk {
		if n != -1 {
			part2 += i * n
		}
	}
	fmt.Println(part2)

	return part1, part2
}
