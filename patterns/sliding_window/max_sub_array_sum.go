package slidingwindow

func MaxSubArraySum(nums []int, n int) int {
	if len(nums) < n {
		return 0
	}

	maximum, sum := 0, 0

	for i := 0; i < n; i++ {
		sum += nums[i]
	}

	maximum = sum

	for i := n; i < len(nums); i++ {
		sum += nums[i] - nums[i-n]
		maximum = max(maximum, sum)
	}

	return maximum
}
