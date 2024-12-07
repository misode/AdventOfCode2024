package utils

import (
	"log"
	"os"
	"path/filepath"
	"regexp"
	"runtime"
	"strconv"
	"strings"
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
