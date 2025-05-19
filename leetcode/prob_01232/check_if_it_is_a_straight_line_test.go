package prob_01232

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestIt(t *testing.T) {
	cases := []struct {
		coordinates [][]int
		expected    bool
	}{
		{coordinates: [][]int{{2, 1}, {4, 2}, {6, 3}}, expected: true},
		{coordinates: [][]int{{2, 4}, {2, 5}, {2, 8}}, expected: true},
		{coordinates: [][]int{{1, 2}, {2, 3}, {3, 4}, {4, 5}, {5, 6}, {6, 7}}, expected: true},
		{coordinates: [][]int{{1, 1}, {2, 1}, {3, 4}, {4, 5}, {5, 6}, {7, 7}}, expected: false},
	}
	for _, c := range cases {
		assert.Equal(t, c.expected, checkStraightLine(c.coordinates))
	}
}
