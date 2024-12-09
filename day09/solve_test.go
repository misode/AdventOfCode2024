package day09

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSolve(t *testing.T) {
	part1, part2 := Solve()
	assert.Equal(t, part1, 6430446922192)
	assert.Equal(t, part2, 6460170593016)
}
