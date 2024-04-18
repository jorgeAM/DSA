package multiplepointer

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCountUniqueValues(t *testing.T) {
	assert.Equal(t, CountUniqueValues([]int{1, 1, 1, 1, 2}), 2)
	assert.Equal(t, CountUniqueValues([]int{1, 2, 3, 4, 4, 4, 7, 7, 12, 12, 13}), 7)
	assert.Equal(t, CountUniqueValues([]int{}), 0)
	assert.Equal(t, CountUniqueValues([]int{1}), 1)
	assert.Equal(t, CountUniqueValues([]int{1, 1, 1}), 1)
	assert.Equal(t, CountUniqueValues([]int{-2, -1, -1, 0, 1}), 4)
}
