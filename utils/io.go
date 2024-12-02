package utils

import (
	"log"
	"os"
	"path/filepath"
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
