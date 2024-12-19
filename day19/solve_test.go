package day19

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSolve(t *testing.T) {
	part1, part2 := Solve()
	assert.Equal(t, part1, 209)
	assert.Equal(t, part2, 777669668613191)
}
