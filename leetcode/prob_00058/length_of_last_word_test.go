package prob_00058

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestIt(t *testing.T) {
	cases := []struct {
		s        string
		expected int
	}{
		{s: "Hello World", expected: 5},
		{s: "   fly me   to   the moon  ", expected: 4},
		{s: "luffy is still joyboy", expected: 6},
		{s: " ", expected: 0},
	}
	for _, c := range cases {
		assert.Equal(t, c.expected, lengthOfLastWord(c.s))
	}
}
