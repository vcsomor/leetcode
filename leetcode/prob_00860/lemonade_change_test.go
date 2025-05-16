package prob_00860

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestIt(t *testing.T) {
	cases := []struct {
		bills    []int
		expected bool
	}{
		{bills: []int{5}, expected: true},
		{bills: []int{10}, expected: false},
		{bills: []int{20}, expected: false},
		{bills: []int{5, 5, 5, 10, 20}, expected: true},
		{bills: []int{5, 5, 10, 10, 20}, expected: false},
	}
	for _, c := range cases {
		assert.Equal(t, c.expected, lemonadeChange(c.bills))
	}
}
