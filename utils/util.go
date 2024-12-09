package utils

import (
	"fmt"
	"log"
	"os"
	"path/filepath"
	"regexp"
	"runtime"
	"strconv"
	"strings"
	"time"
	"unicode/utf8"
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
		lines = append(lines, line)
	}

	if len(lines[len(lines)-1]) == 0 {
		lines = lines[:len(lines)-1]
	}

	return lines
}

type DayTimer struct {
	start *time.Time
	parse *time.Time
	part1 *time.Time
	part2 *time.Time
}

func (t *DayTimer) getPrev() time.Time {
	if t.part2 != nil {
		return *t.part2
	} else if t.part1 != nil {
		return *t.part1
	} else if t.parse != nil {
		return *t.parse
	} else if t.start != nil {
		return *t.start
	} else {
		return time.Now()
	}
}

func StartDay(day int) DayTimer {
	fmt.Printf("=== Day %02d ===\n", day)
	now := time.Now()
	return DayTimer{start: &now, parse: nil, part1: nil, part2: nil}
}

func (t *DayTimer) Parsed() {
	now := time.Now()
	t.parse = &now
}

func (t *DayTimer) Part1(val any) {
	now := time.Now()
	diff := now.Sub(t.getPrev())
	t.part1 = &now
	fmt.Printf(" %v\t\t(%v)\n", val, diff)
}

func (t *DayTimer) Part2(val any) {
	now := time.Now()
	diff := now.Sub(t.getPrev())
	t.part2 = &now
	fmt.Printf(" %v\t\t(%v)\n", val, diff)
}

func (t *DayTimer) Parts(val1 any, val2 any) {
	now := time.Now()
	diff := now.Sub(t.getPrev())
	t.part1 = &now
	t.part2 = &now
	fmt.Printf(" %v\n", val1)
	fmt.Printf(" %v\t\t(%v)\n", val2, diff)
}

func SplitLinesOnEmpty(lines []string) [][]string {
	groups := make([][]string, 0)
	cur := make([]string, 0)
	for _, line := range lines {
		if len(line) == 0 {
			if len(cur) > 0 {
				groups = append(groups, cur)
			}
			cur = make([]string, 0)
		} else {
			cur = append(cur, line)
		}
	}
	if len(cur) > 0 {
		groups = append(groups, cur)
	}
	return groups
}

func SplitInts(source string, sep string) []int {
	parts := strings.Split(source, sep)
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

func ReverseStr(s string) string {
	size := len(s)
	buf := make([]byte, size)
	for start := 0; start < size; {
		r, n := utf8.DecodeRuneInString(s[start:])
		start += n
		utf8.EncodeRune(buf[size-start:], r)
	}
	return string(buf)
}

func SliceEqual[T comparable](a []T, b []T) bool {
	if len(a) != len(b) {
		return false
	}
	for i, val := range a {
		if val != b[i] {
			return false
		}
	}
	return true
}

func ForCombinations[T any](list []T, fn func(a T, b T)) {
	for i := 0; i < len(list); i++ {
		for j := i + 1; j < len(list); j++ {
			fn(list[i], list[j])
		}
	}
}
