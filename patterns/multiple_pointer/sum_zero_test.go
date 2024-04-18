package multiplepointer

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSumZero(t *testing.T) {
	assert.Equal(t, SumZero([]int{-3, -2, -1, 0, 1, 2, 3}), []int{-3, 3})
	assert.Equal(t, SumZero([]int{-2, 0, 1, 3}), []int{0})
	assert.Equal(t, SumZero([]int{1, 2, 3}), []int{0})
}
