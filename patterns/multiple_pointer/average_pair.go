package multiplepointer

func AveragePair(nums []int, average float64) bool {
	if len(nums) == 0 {
		return false
	}

	l, r := 0, len(nums)-1

	for l < r {
		avg := (float64(nums[l]) + float64(nums[r])) / 2

		if avg == average {
			return true
		}

		if avg > average {
			r--
		} else {
			l++
		}
	}

	return false
}
