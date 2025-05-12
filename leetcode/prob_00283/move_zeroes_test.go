package prob_00283

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestIt(t *testing.T) {
	cases := []struct {
		input    []int
		expected []int
	}{
		{input: []int{0}, expected: []int{0}},
		{input: []int{1, 0, 1}, expected: []int{1, 1, 0}},
		{input: []int{1, 3, 12}, expected: []int{1, 3, 12}},
		{input: []int{0, 1, 0, 3, 12}, expected: []int{1, 3, 12, 0, 0}},
		{input: []int{0, 0, 0, 1, 0, 0, 3, 12, 15}, expected: []int{1, 3, 12, 15, 0, 0, 0, 0, 0}},
		{input: []int{0, 0, 0, 1, 0, 0, 3, 12, 15, 0, 0}, expected: []int{1, 3, 12, 15, 0, 0, 0, 0, 0, 0, 0}},
	}
	for _, c := range cases {
		data := c.input
		moveZeroes(data)
		assert.Equal(t, c.expected, data)
	}
}
