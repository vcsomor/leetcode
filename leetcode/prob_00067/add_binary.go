package prob_00067

import (
	"strings"
)

func addBinary(a string, b string) string {
	num0, num1 := a, b
	if len(a) < len(b) {
		num0, num1 = b, a
	}

	i0 := len(num0) - 1 // longer number
	i1 := len(num1) - 1
	result := strings.Builder{}
	result.Grow(len(num0) + 1)

	carryOver := false

	for i0 >= 0 {
		co, res := doAdd(isOne(num0, i0), isOne(num1, i1), carryOver)
		carryOver = co

		result.WriteByte(res)

		i0--
		i1--
	}

	res := reverse(result.String())
	if carryOver {
		return "1" + res
	}
	return res
}

func reverse(s string) string {
	r := strings.Builder{}
	r.Grow(len(s))
	for i := len(s) - 1; i >= 0; i-- {
		r.WriteByte(s[i])
	}
	return r.String()
}

func doAdd(a, b, carryOver bool) (bool, byte) {
	if a && b && carryOver {
		return true, '1'
	}
	if (a && b) || (a && carryOver) || (b && carryOver) {
		return true, '0'
	}
	if carryOver {
		return false, '1'
	}
	if a || b {
		return false, '1'
	}
	return false, '0'
}

func isOne(num string, i int) bool {
	if i < 0 {
		return false
	}
	return num[i] == '1'
}
