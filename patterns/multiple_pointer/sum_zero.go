package multiplepointer

func SumZero(nums []int) []int {
	l, r := 0, len(nums)-1

	for l < r {
		sum := nums[l] + nums[r]

		if sum == 0 {
			return []int{nums[l], nums[r]}
		} else if sum > 0 {
			r--
		} else {
			l++
		}
	}

	return []int{0}
}
