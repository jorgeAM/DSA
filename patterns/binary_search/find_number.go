package binarysearch

func FindNumber(nums []int, num int) int {
	start, end := 0, len(nums)-1

	for start <= end {
		middle := (end + start) / 2

		if nums[middle] == num {
			return middle
		} else if nums[middle] > num {
			end = middle - 1
		} else {
			start = middle + 1
		}
	}

	return -1
}
