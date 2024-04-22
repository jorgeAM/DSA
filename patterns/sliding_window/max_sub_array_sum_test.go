package slidingwindow

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestMasSubArraySum(t *testing.T) {
	assert.Equal(t, MaxSubArraySum([]int{1, 2, 5, 2, 8, 1, 5}, 2), 10)
	assert.Equal(t, MaxSubArraySum([]int{1, 2, 5, 2, 8, 1, 5}, 4), 17)
	assert.Equal(t, MaxSubArraySum([]int{4, 2, 1, 6, 2}, 4), 13)
	assert.Equal(t, MaxSubArraySum([]int{}, 4), 0)
	assert.Equal(t, MaxSubArraySum([]int{4, 2, 1, 6}, 1), 6)
}
