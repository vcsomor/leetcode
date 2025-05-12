package prob_00066

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestIt(t *testing.T) {
	cases := []struct {
		digits   []int
		expected []int
	}{
		{digits: []int{1, 2, 3}, expected: []int{1, 2, 4}},
		{digits: []int{4, 3, 2, 1}, expected: []int{4, 3, 2, 2}},
		{digits: []int{9, 9, 9, 9}, expected: []int{1, 0, 0, 0, 0}},
		{digits: []int{1, 2, 1, 9}, expected: []int{1, 2, 2, 0}},
	}
	for _, c := range cases {
		assert.Equal(t, c.expected, plusOne(c.digits))
	}
}
