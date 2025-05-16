package prob_00976

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestIt(t *testing.T) {
	cases := []struct {
		nums     []int
		expected int
	}{
		{nums: []int{2, 1, 2}, expected: 5},
		{nums: []int{1, 2, 1, 10}, expected: 0},
	}
	for _, c := range cases {
		assert.Equal(t, c.expected, largestPerimeter(c.nums))
	}
}
