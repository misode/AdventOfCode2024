package day17

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSolve(t *testing.T) {
	part1, part2 := Solve()
	assert.Equal(t, part1, "3,1,4,3,1,7,1,6,3")
	assert.Equal(t, part2, 0)
}
