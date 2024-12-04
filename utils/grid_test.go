package utils

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGridMatch(t *testing.T) {
	grid := MakeGrid([]string{"ABC", "DEF"})
	assert.Equal(t, grid.Match(0, 0, "AD"), false)
	assert.Equal(t, grid.Match(0, 0, "A,D"), true)
	assert.Equal(t, grid.Match(0, 0, "EF"), false)
	assert.Equal(t, grid.Match(1, 1, "EF"), true)
}

func TestGridMatchSym(t *testing.T) {
	grid := MakeGrid([]string{"ABC", "DEF"})
	assert.Equal(t, grid.MatchSym(0, 1, "AD"), false)
	assert.Equal(t, grid.MatchSym(0, 1, "BC"), true)
	assert.Equal(t, grid.MatchSym(0, 1, "CB"), true)
}

func TestGridTranspose(t *testing.T) {
	grid := MakeGrid([]string{"ABC", "DEF"})
	transposed := grid.Transpose()
	lines := transposed.ToLines()
	assert.Equal(t, lines, []string{"AD", "BE", "CF"})
}
