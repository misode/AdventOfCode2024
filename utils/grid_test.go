package utils

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGridTranspose(t *testing.T) {
	grid := MakeGrid([]string{"ABC", "DEF"})
	transposed := grid.Transpose()
	lines := transposed.ToLines()
	assert.Equal(t, lines, []string{"AD", "BE", "CF"})
}

func TestRuneToLines(t *testing.T) {
	grid := MakeGrid([]string{"ABC", "DEF"})
	lines := grid.ToLines()
	assert.Equal(t, lines, []string{"ABC", "DEF"})
}

func TestIntToLines(t *testing.T) {
	grid := MakeIntGrid([]string{"012", "345"})
	lines := grid.ToLines()
	assert.Equal(t, lines, []string{"012", "345"})
}
