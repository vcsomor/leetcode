// Copyright (c) 2017-2025, Cloudera, Inc. All rights reserved.
//
// Except as expressly permitted in a written agreement between you or your company
// and Cloudera, Inc. or an authorized affiliate or partner thereof, any use,
// reproduction, modification, redistribution, sharing, lending or other exploitation
// of all or any part of the contents of this software is strictly prohibited.

package prob_00283

func moveZeroes(nums []int) {
	moveZeroesSimple(nums)
}

func moveZeroesSimple(nums []int) {
	if len(nums) == 1 {
		return
	}

	left := 0
	right := 0

	for right < len(nums) {
		if nums[right] != 0 {
			swap(nums, left, right)
			left++
		}
		right++
	}
}

func moveZeroesNaive(nums []int) {
	if len(nums) == 1 {
		return
	}

	readPtr := 0
	writePtr := seekWritePtr(nums, len(nums)-1)

	for readPtr < writePtr {
		zeroBlockStart, zeroBlockLength := findZeroBlock(nums, readPtr)

		if zeroBlockStart == -1 {
			break
		}

		shiftBlockLeft(nums, zeroBlockStart, writePtr, zeroBlockLength)
		readPtr++
		writePtr -= zeroBlockLength
	}
}

func shiftBlockLeft(nums []int, from int, to int, shift int) {
	for i := from; (i + shift) <= to; i++ {
		swap(nums, i, i+shift)
	}
}

func swap(nums []int, at int, to int) {
	x := nums[at]
	nums[at] = nums[to]
	nums[to] = x
}

func seekWritePtr(num []int, w int) int {
	for i := w; i >= 0; i-- {
		if num[i] != 0 {
			return i
		}
	}
	return 0
}

func findZeroBlock(nums []int, ptr int) (int, int) {
	blkLen := 0

	foundAt := -1
	readPtr := ptr
	for readPtr < len(nums) {
		if foundAt == -1 {
			if nums[readPtr] == 0 {
				foundAt = readPtr
				blkLen = 1
			}
		} else if nums[readPtr] == 0 {
			blkLen++
		} else {
			break
		}
		readPtr++
	}
	return foundAt, blkLen
}
