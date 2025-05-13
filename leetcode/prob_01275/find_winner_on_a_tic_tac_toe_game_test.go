package prob_01275

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestIt(t *testing.T) {
	cases := []struct {
		moves    [][]int
		expected string
	}{
		{moves: [][]int{{0, 0}, {2, 0}, {1, 1}, {2, 1}, {2, 2}}, expected: "A"},
		{moves: [][]int{{0, 0}, {1, 1}, {0, 1}, {0, 2}, {1, 0}, {2, 0}}, expected: "B"},
		{moves: [][]int{{0, 0}, {1, 1}, {2, 0}, {1, 0}, {1, 2}, {2, 1}, {0, 1}, {0, 2}, {2, 2}}, expected: "Draw"},
		{moves: [][]int{{0, 0}, {1, 1}}, expected: "Pending"},
	}
	for _, c := range cases {
		assert.Equal(t, c.expected, tictactoe(c.moves))
	}
}
