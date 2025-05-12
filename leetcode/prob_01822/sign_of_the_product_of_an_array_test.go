package prob_01822

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestIt(t *testing.T) {
	cases := []struct {
		digits   []int
		expected int
	}{
		{digits: []int{-1, -2, -3, -4, 3, 2, 1}, expected: 1},
		{digits: []int{1, 5, 0, 2, -3}, expected: 0},
		{digits: []int{-1, 1, -1, 1, -1}, expected: -1},
	}
	for _, c := range cases {
		assert.Equal(t, c.expected, arraySign(c.digits))
	}
}
