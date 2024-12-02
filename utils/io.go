package utils

import (
	"log"
	"os"
	"strings"
)

func ReadInput(name string) []string {
	content, err := os.ReadFile(name)
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
