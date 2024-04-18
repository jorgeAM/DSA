package multiplepointer

import "fmt"

func CountUniqueValues(nums []int) int {

	/**
		// time complexity: O(n)
		// space complexity: O(n)

		if len(nums) == 0 {
			return 0
		}

		counter := make(map[int]int)

		for i := range nums {
			counter[nums[i]]++
		}

		return len(counter)

	**/

	// time complexity: O(n)
	// space complexity: O(1)

	if len(nums) == 0 {
		return 0
	}

	counter, l := 1, 0
	for r := 1; r < len(nums); r++ {
		if nums[l] == nums[r] {
			continue
		}

		l = r
		counter++
	}

	fmt.Println(l)

	return counter
}
