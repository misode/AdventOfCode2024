package utils

import (
	"log"
	"os"
	"path/filepath"
	"regexp"
	"runtime"
	"strconv"
	"strings"
)

func ReadInput(name string) []string {
	_, currentFile, _, _ := runtime.Caller(1)
	currentDir := filepath.Dir(currentFile)
	inputFilePath := filepath.Join(currentDir, name)

	content, err := os.ReadFile(inputFilePath)
	if err != nil {
		log.Fatal(err)
	}

	var lines []string

	for _, line := range strings.Split(string(content), "\n") {
		line, _ := strings.CutSuffix(line, "\r")
		if len(line) == 0 {
			continue
		}
		lines = append(lines, line)
	}

	return lines
}

func SplitInts(source string) []int {
	parts := strings.Split(source, " ")
	ints := make([]int, len(parts))
	i := 0
	for _, part := range parts {
		num, err := strconv.Atoi(part)
		if err == nil {
			ints[i] = num
			i += 1
		}
	}
	return ints[:i]
}

func Counter[K comparable](values []K) map[K]int {
	counts := make(map[K]int)
	for _, val := range values {
		counts[val] = counts[val] + 1
	}
	return counts
}

func StrToInt(value string) int {
	num, err := strconv.Atoi(value)
	if err != nil {
		log.Fatal(err)
	}
	return num
}

func FindMatches(pattern string, source string) [][]string {
	regex, err := regexp.Compile(pattern)
	if err != nil {
		log.Fatal(err)
	}
	return regex.FindAllStringSubmatch(source, -1)
}

func MakeGrid(lines []string) [][]rune {
	grid := make([][]rune, 0, len(lines))
	for _, line := range lines {
		grid = append(grid, []rune(line))
	}
	return grid
}

func AssertEqual[T comparable](a T, b T) {
	if a != b {
		log.Fatal("AssertEqual")
	}
}

func MatchSubGrid(grid [][]rune, r int, c int, search string) bool {
	subgrid := MakeGrid(strings.Split(search, ","))

	if r+len(subgrid) > len(grid) || c+len(subgrid[0]) > len(grid[0]) {
		return false
	}
	for dr := 0; dr < len(subgrid); dr++ {
		for dc := 0; dc < len(subgrid[0]); dc++ {
			if subgrid[dr][dc] != ' ' && grid[r+dr][c+dc] != subgrid[dr][dc] {
				return false
			}
		}
	}
	return true
}
