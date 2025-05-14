package prob_01672

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestIt(t *testing.T) {
	cases := []struct {
		accounts [][]int
		expected int
	}{
		{accounts: [][]int{{1, 2, 3}, {3, 2, 1}}, expected: 6},
		{accounts: [][]int{{1, 5}, {7, 3}, {3, 5}}, expected: 10},
	}
	for _, c := range cases {
		assert.Equal(t, c.expected, maximumWealth(c.accounts))
	}
}
