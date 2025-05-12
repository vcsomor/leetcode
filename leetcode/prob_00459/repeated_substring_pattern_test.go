package prob_00459

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestIt(t *testing.T) {
	cases := []struct {
		s        string
		expected bool
	}{
		{s: "aabaaba", expected: false},

		{s: "a", expected: false},
		{s: "aa", expected: true},
		{s: "aaa", expected: true},
		{s: "aba", expected: false},
		{s: "abab", expected: true},
		{s: "abcab", expected: false},
		{s: "abcabc", expected: true},
		{s: "abcxabc", expected: false},
		{s: "abcxabcx", expected: true},
		{s: "abababab", expected: true},
		{s: "abcdabcd", expected: true},
	}
	for _, c := range cases {
		assert.Equal(t, c.expected, repeatedSubstringPattern(c.s))
	}
}
