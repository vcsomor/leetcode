package prob_01491

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestIt(t *testing.T) {
	cases := []struct {
		salary   []int
		expected float64
	}{
		{salary: []int{4000, 3000, 1000, 2000}, expected: 2500.},
		{salary: []int{1000, 2000, 3000}, expected: 2000.},
	}
	for _, c := range cases {
		assert.Equal(t, c.expected, average(c.salary))
	}
}
