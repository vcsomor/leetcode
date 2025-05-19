package prob_00043

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestIt(t *testing.T) {
	cases := []struct {
		num1, num2 string
		expected   string
	}{
		{num1: "2", num2: "3", expected: "6"},
		{num1: "123", num2: "456", expected: "56088"},
	}
	for _, c := range cases {
		assert.Equal(t, c.expected, multiply(c.num1, c.num2))
	}
}
