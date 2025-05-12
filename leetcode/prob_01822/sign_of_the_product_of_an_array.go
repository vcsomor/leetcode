// Copyright (c) 2017-2025, Cloudera, Inc. All rights reserved.
//
// Except as expressly permitted in a written agreement between you or your company
// and Cloudera, Inc. or an authorized affiliate or partner thereof, any use,
// reproduction, modification, redistribution, sharing, lending or other exploitation
// of all or any part of the contents of this software is strictly prohibited.

package prob_01822

func arraySign(nums []int) int {
	minusSign := false

	for _, n := range nums {
		if n == 0 {
			return 0
		}
		if n < 0 {
			minusSign = !minusSign
		}
	}

	if minusSign {
		return -1
	}
	return 1
}
