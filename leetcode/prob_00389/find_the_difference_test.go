// Copyright (c) 2017-2025, Cloudera, Inc. All rights reserved.
//
// Except as expressly permitted in a written agreement between you or your company
// and Cloudera, Inc. or an authorized affiliate or partner thereof, any use,
// reproduction, modification, redistribution, sharing, lending or other exploitation
// of all or any part of the contents of this software is strictly prohibited.

package prob_00389

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestIt(t *testing.T) {
	cases := []struct {
		s        string
		t        string
		expected byte
	}{
		{s: "", t: "y", expected: 'y'},
		{s: "abcda", t: "aebcda", expected: 'e'},
	}
	for _, c := range cases {
		assert.Equal(t, c.expected, findTheDifference(c.s, c.t))
	}
}
