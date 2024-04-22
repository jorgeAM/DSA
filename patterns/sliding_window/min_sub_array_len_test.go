package slidingwindow

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestMinArrayLen(t *testing.T) {
	assert.Equal(t, MinSubArrayLen([]int{2, 3, 1, 2, 4, 3}, 7), 2)
	assert.Equal(t, MinSubArrayLen([]int{2, 1, 6, 5, 4}, 9), 2)
	assert.Equal(t, MinSubArrayLen([]int{3, 1, 7, 11, 2, 9, 8, 21, 62, 33, 19}, 52), 1)
	assert.Equal(t, MinSubArrayLen([]int{1, 4, 16, 22, 5, 7, 8, 9, 10}, 39), 3)
	assert.Equal(t, MinSubArrayLen([]int{1, 4, 16, 22, 5, 7, 8, 9, 10}, 55), 5)
	assert.Equal(t, MinSubArrayLen([]int{4, 3, 3, 8, 1, 2, 3}, 11), 2)
}
