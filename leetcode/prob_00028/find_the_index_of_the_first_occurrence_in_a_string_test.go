// Copyright (c) 2017-2025, Cloudera, Inc. All rights reserved.
//
// Except as expressly permitted in a written agreement between you or your company
// and Cloudera, Inc. or an authorized affiliate or partner thereof, any use,
// reproduction, modification, redistribution, sharing, lending or other exploitation
// of all or any part of the contents of this software is strictly prohibited.

package prob_00028

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestIt(t *testing.T) {
	cases := []struct {
		haystack string
		needle   string
		expected int
	}{
		{haystack: "a", needle: "a", expected: 0},
		{haystack: "aaa", needle: "a", expected: 0},
		{haystack: "aaba", needle: "ab", expected: 1},
		{haystack: "aaba", needle: "ab", expected: 1},
		{haystack: "aa", needle: "ab", expected: -1},
		{haystack: "a", needle: "ab", expected: -1},
		{haystack: "aaaaaaaaaaaaaaaaabcxxxx", needle: "abc", expected: 16},
		{haystack: "aaaaaaaaaaaaaaaaabexxx", needle: "abc", expected: -1},
	}
	for _, c := range cases {
		assert.Equal(t, c.expected, strStr(c.haystack, c.needle))
	}
}
