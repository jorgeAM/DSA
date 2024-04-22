package multiplepointer

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestIsSubsequence(t *testing.T) {
	assert.True(t, IsSubsequence("hello", "hello world"))
	assert.True(t, IsSubsequence("sing", "sting"))
	assert.True(t, IsSubsequence("abc", "abracadabra"))
	assert.False(t, IsSubsequence("abc", "acb"))
}
