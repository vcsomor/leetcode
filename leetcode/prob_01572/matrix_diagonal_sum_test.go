package prob_01572

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestIt(t *testing.T) {
	cases := []struct {
		matrix   [][]int
		expected int
	}{
		{matrix: [][]int{
			{1, 2, 3},
			{4, 5, 6},
			{7, 8, 9},
		},
			expected: 0},

		{matrix: [][]int{
			{1, 1, 1, 1},
			{1, 1, 1, 1},
			{1, 1, 1, 1},
			{1, 1, 1, 1},
		},
			expected: 8},

		{matrix: [][]int{
			{5},
		},
			expected: 5},
	}
	for _, c := range cases {
		assert.Equal(t, c.expected, diagonalSum(c.matrix))
	}
}
