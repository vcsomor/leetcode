package prob_01502

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestIt(t *testing.T) {
	cases := []struct {
		arr      []int
		expected bool
	}{
		{arr: []int{3, 5}, expected: true},
		{arr: []int{3, 5, 1}, expected: true},
		{arr: []int{1, 2, 4}, expected: false},
		{arr: []int{3, 1, 1, 2, 3}, expected: false},
		{arr: []int{0, 0, 0, 0}, expected: true},
	}
	for _, c := range cases {
		assert.Equal(t, c.expected, canMakeArithmeticProgression(c.arr))
	}
}
