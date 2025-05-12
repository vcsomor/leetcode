// Copyright (c) 2017-2025, Cloudera, Inc. All rights reserved.
//
// Except as expressly permitted in a written agreement between you or your company
// and Cloudera, Inc. or an authorized affiliate or partner thereof, any use,
// reproduction, modification, redistribution, sharing, lending or other exploitation
// of all or any part of the contents of this software is strictly prohibited.

package prob_01822

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestIt(t *testing.T) {
	cases := []struct {
		digits   []int
		expected int
	}{
		{digits: []int{-1, -2, -3, -4, 3, 2, 1}, expected: 1},
		{digits: []int{1, 5, 0, 2, -3}, expected: 0},
		{digits: []int{-1, 1, -1, 1, -1}, expected: -1},
	}
	for _, c := range cases {
		assert.Equal(t, c.expected, arraySign(c.digits))
	}
}
