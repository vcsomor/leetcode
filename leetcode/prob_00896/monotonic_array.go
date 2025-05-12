// Copyright (c) 2017-2025, Cloudera, Inc. All rights reserved.
//
// Except as expressly permitted in a written agreement between you or your company
// and Cloudera, Inc. or an authorized affiliate or partner thereof, any use,
// reproduction, modification, redistribution, sharing, lending or other exploitation
// of all or any part of the contents of this software is strictly prohibited.

package prob_00869

const (
	decreasing int = -1
	stay       int = 0
	increasing int = 1
)

func isMonotonic(nums []int) bool {
	lastTendency := stay
	for i := 2; i < len(nums); i++ {
		if lastTendency == stay {
			lastTendency = tendency(nums[i-2], nums[i-1])
		}
		currTendency := tendency(nums[i-1], nums[i])
		if currTendency == stay {
			continue
		}

		if lastTendency == stay {
			lastTendency = currTendency
			continue
		}

		if currTendency != lastTendency {
			return false
		}
	}
	return true
}

func tendency(n0, n1 int) int {
	if n0 < n1 {
		return increasing
	}
	if n0 > n1 {
		return decreasing
	}
	return stay
}
