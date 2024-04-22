package multiplepointer

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestAveragePair(t *testing.T) {
	assert.True(t, AveragePair([]int{1, 2, 3}, 2.5))
	assert.True(t, AveragePair([]int{1, 3, 3, 5, 6, 7, 10, 12, 19}, 8))
	assert.False(t, AveragePair([]int{-1, 0, 3, 4, 5, 6}, 4.1))
	assert.False(t, AveragePair([]int{}, 4))
}
