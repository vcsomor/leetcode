package prob_template

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestIt(t *testing.T) {
	cases := []struct {
		input    string
		expected bool
	}{
		{input: "UD", expected: true},
		{input: "LL", expected: false},
	}
	for _, c := range cases {
		assert.Equal(t, c.expected, judgeCircle(c.input))
	}
}
