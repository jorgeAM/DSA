package frequencycounter

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestValidAnagrams(t *testing.T) {
	assert.True(t, ValidAnagram("", ""))
	assert.False(t, ValidAnagram("aaz", "zza"))
	assert.True(t, ValidAnagram("anagram", "nagaram"))
	assert.False(t, ValidAnagram("rat", "car"))
	assert.False(t, ValidAnagram("awesome", "awesom"))
	assert.True(t, ValidAnagram("qwerty", "qeywrt"))
	assert.True(t, ValidAnagram("cinema", "iceman"))
}
