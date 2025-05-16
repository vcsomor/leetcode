package prob_00976

import "sort"

func largestPerimeter(nums []int) int {
	sort.Ints(nums)

	for c := len(nums) - 1; c >= 2; c-- {
		a := c - 1
		b := c - 2
		if (nums[b] + nums[a]) > nums[c] {
			return nums[c] + nums[b] + nums[a]
		}
	}
	return 0
}
