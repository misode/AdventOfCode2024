package day11

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSolve(t *testing.T) {
	part1, part2 := Solve()
	assert.Equal(t, part1, 233050)
	assert.Equal(t, part2, 276661131175807)
}
