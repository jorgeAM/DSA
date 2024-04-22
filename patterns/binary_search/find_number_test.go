package binarysearch

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestFindNumber(t *testing.T) {
	assert.Equal(t, FindNumber([]int{1, 2, 3, 4, 5}, 2), 1)
	assert.Equal(t, FindNumber([]int{1, 2, 3, 4, 5}, 3), 2)
	assert.Equal(t, FindNumber([]int{1, 2, 3, 4, 5}, 5), 4)
	assert.Equal(t, FindNumber([]int{1, 2, 3, 4, 5}, 6), -1)
	assert.Equal(t, FindNumber([]int{5, 6, 10, 13, 14, 18, 30, 34, 35, 37, 40, 44, 64, 79, 84, 86, 95, 96, 98, 99}, 95), 16)
	assert.Equal(t, FindNumber([]int{5, 6, 10, 13, 14, 18, 30, 34, 35, 37, 40, 44, 64, 79, 84, 86, 95, 96, 98, 99}, 100), -1)
}
