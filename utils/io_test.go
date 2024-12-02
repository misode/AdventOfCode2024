package utils

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSplitInts(t *testing.T) {
	assert.Equal(t, SplitInts("1 4 2"), []int{1, 4, 2})
	assert.Equal(t, SplitInts("8 test 3"), []int{8, 3})
}

func TestCounter(t *testing.T) {
	assert.Equal(t, Counter([]int{2, 4, 3, 3}), map[int]int{2: 1, 3: 2, 4: 1})
}
