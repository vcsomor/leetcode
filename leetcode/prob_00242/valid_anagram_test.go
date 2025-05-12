package prob_00028

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestIt(t *testing.T) {
	cases := []struct {
		s        string
		t        string
		expected bool
	}{
		{s: "anagram", t: "nagaram", expected: true},
		{s: "x", t: "x", expected: true},
		{s: "car", t: "rat", expected: false},
		{s: "x", t: "y", expected: false},
	}
	for _, c := range cases {
		assert.Equal(t, c.expected, isAnagram(c.s, c.t))
	}
}
