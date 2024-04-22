package frequencycounter

import (
	"strconv"
)

func SameFrequency(a, b int) bool {
	stringA := strconv.FormatInt(int64(a), 10)
	stringB := strconv.FormatInt(int64(b), 10)

	if len(stringA) != len(stringB) {
		return false
	}

	counterA := make(map[int]int)
	counterB := make(map[int]int)

	for a > 0 && b > 0 {
		counterA[a%10]++
		counterB[b%10]++

		a /= 10
		b /= 10
	}

	for k := range counterA {
		if counterA[k] != counterB[k] {
			return false
		}
	}

	return true
}
