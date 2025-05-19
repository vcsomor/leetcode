package prob_00067

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestIt(t *testing.T) {
	cases := []struct {
		a, b     string
		expected string
	}{
		{a: "1", b: "101", expected: "110"},
		{a: "1010", b: "1011", expected: "10101"},

		{a: "11", b: "1", expected: "100"},
		{a: "100", b: "0", expected: "100"},
	}
	for _, c := range cases {
		assert.Equal(t, c.expected, addBinary(c.a, c.b))
	}
}
