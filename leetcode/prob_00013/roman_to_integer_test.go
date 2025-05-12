package prob_00013

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestIt(t *testing.T) {
	cases := []struct {
		roman    string
		expected int
	}{
		{roman: "I", expected: 1},
		{roman: "LVIII", expected: 58},
		{roman: "MCMXCIV", expected: 1994},

		{roman: "I", expected: 1},
		{roman: "V", expected: 5},
		{roman: "X", expected: 10},
		{roman: "L", expected: 50},
		{roman: "C", expected: 100},
		{roman: "D", expected: 500},
		{roman: "M", expected: 1000},

		{roman: "IV", expected: 4},
		{roman: "IX", expected: 9},
		{roman: "XL", expected: 40},
		{roman: "XC", expected: 90},
		{roman: "CD", expected: 400},
		{roman: "CM", expected: 900},
	}
	for _, c := range cases {
		assert.Equal(t, c.expected, romanToInt(c.roman))
	}
}
