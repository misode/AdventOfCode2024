package day07

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSolve(t *testing.T) {
	part1, part2 := Solve()
	assert.Equal(t, part1, 2501605301465)
	assert.Equal(t, part2, 44841372855953)
}
