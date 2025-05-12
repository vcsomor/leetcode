package prob_00682

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestIt(t *testing.T) {
	cases := []struct {
		ops      []string
		expected int
	}{
		{ops: []string{"5", "2", "C", "D", "+"}, expected: 30},
		{ops: []string{"5", "-2", "4", "C", "D", "9", "+", "+"}, expected: 27},
		{ops: []string{"1", "C"}, expected: 0},
		{ops: []string{"0"}, expected: 0},
	}
	for _, c := range cases {
		assert.Equal(t, c.expected, calPoints(c.ops))
	}
}
