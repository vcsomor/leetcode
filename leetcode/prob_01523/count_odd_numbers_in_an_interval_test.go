package prob_01523

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestIt(t *testing.T) {
	cases := []struct {
		low, high int
		expected  int
	}{
		{low: 3, high: 7, expected: 3},
		{low: 8, high: 10, expected: 1},
	}
	for _, c := range cases {
		assert.Equal(t, c.expected, countOdds(c.low, c.high))
	}
}
