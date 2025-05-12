// Copyright (c) 2017-2025, Cloudera, Inc. All rights reserved.
//
// Except as expressly permitted in a written agreement between you or your company
// and Cloudera, Inc. or an authorized affiliate or partner thereof, any use,
// reproduction, modification, redistribution, sharing, lending or other exploitation
// of all or any part of the contents of this software is strictly prohibited.

package prob_00869

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestIt(t *testing.T) {
	cases := []struct {
		arr      []int
		expected bool
	}{
		{arr: []int{1}, expected: true},
		{arr: []int{1, 1}, expected: true},
		{arr: []int{1, 2}, expected: true},
		{arr: []int{1, 0}, expected: true},
		{arr: []int{1, 1, 1, 1, 3, -2}, expected: false},
		{arr: []int{1, 2, 2, 3}, expected: true},
		{arr: []int{6, 5, 4, 4}, expected: true},
		{arr: []int{1, 3, 2}, expected: false},
	}
	for _, c := range cases {
		assert.Equal(t, c.expected, isMonotonic(c.arr))
	}
}
