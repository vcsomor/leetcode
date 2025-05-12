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
