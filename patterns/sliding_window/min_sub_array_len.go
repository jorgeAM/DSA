package slidingwindow

func MinSubArrayLen(nums []int, num int) int {
	sum := 0
	minLen := len(nums)
	l, r := 0, 0

	for l < len(nums) {
		if sum < num && r < len(nums) {
			sum += nums[r]
			r++
		} else if sum >= num {
			minLen = min(minLen, r-l)
			sum -= nums[l]
			l++
		} else {
			break
		}
	}

	return minLen
}
