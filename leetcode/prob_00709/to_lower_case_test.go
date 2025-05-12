package prob_00709

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestIt(t *testing.T) {
	cases := []struct {
		s        string
		expected string
	}{
		{s: "Hello", expected: "hello"},
	}
	for _, c := range cases {
		assert.Equal(t, c.expected, toLowerCase(c.s))
	}
}
