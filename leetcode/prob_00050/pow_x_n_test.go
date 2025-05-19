package prob_00050

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestIt(t *testing.T) {
	cases := []struct {
		x        float64
		n        int
		expected float64
	}{
		{x: 0.25519, n: -6, expected: 3620.9129870675392},
		{x: 2., n: 10, expected: 1024.},
		{x: 2.1, n: 3, expected: 9.261000000000001},
		{x: 2.1, n: 5, expected: 40.841010000000004},
		{x: 2., n: -2, expected: 0.25},
	}
	for _, c := range cases {
		assert.Equal(t, c.expected, myPow(c.x, c.n))
	}
}
