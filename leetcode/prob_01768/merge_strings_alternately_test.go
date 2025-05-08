// Copyright (c) 2017-2025, Cloudera, Inc. All rights reserved.
//
// Except as expressly permitted in a written agreement between you or your company
// and Cloudera, Inc. or an authorized affiliate or partner thereof, any use,
// reproduction, modification, redistribution, sharing, lending or other exploitation
// of all or any part of the contents of this software is strictly prohibited.

package prob_01768

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestIt(t *testing.T) {
	cases := []struct {
		word1    string
		word2    string
		expected string
	}{
		{},
		{
			word1: "abcdf", expected: "abcdf",
		},
		{
			word2: "efg", expected: "efg",
		},
		{
			word1: "abc", word2: "pqr", expected: "apbqcr",
		},
		{
			word1: "ab", word2: "pqrs", expected: "apbqrs",
		},
		{
			word1: "abcd", word2: "pq", expected: "apbqcd",
		},
	}
	for _, c := range cases {
		assert.Equal(t, c.expected, mergeAlternately(c.word1, c.word2))
	}
}
