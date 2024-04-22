package frequencycounter

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSameFrequency(t *testing.T) {
	assert.True(t, SameFrequency(182, 281))
	assert.False(t, SameFrequency(34, 14))
	assert.True(t, SameFrequency(3589578, 5879385))
	assert.False(t, SameFrequency(22, 222))
}
