package prob_template

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestIt(t *testing.T) {
	cases := []struct {
		input    string
		expected int
	}{
		{input: "", expected: 0},
	}
	for _, c := range cases {
		assert.Equal(t, c.expected, test(c.input))
	}
}
