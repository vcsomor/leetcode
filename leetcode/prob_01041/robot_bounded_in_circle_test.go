package prob_01041

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestIt(t *testing.T) {
	cases := []struct {
		input    string
		expected bool
	}{
		{input: "GGLLGG", expected: true},
		{input: "GG", expected: false},
		{input: "GL", expected: true},
	}
	for _, c := range cases {
		assert.Equal(t, c.expected, isRobotBounded(c.input))
	}
}
