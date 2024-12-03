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
