package prob_00043

import "strings"

func multiply(num1 string, num2 string) string {
	if num1 == "0" || num2 == "0" {
		return "0"
	}

	n1 := reverse(num1)
	l := len(num2)

	var items [][]int
	for j := 0; j < l; j++ {
		items = append(items, singleMultiply(addLeadingZeros(n1, l-j-1), num2[j]))
	}

	return toString(sum(items))
}

func addLeadingZeros(n1 string, i int) string {
	if i == 0 {
		return n1
	}
	b := strings.Builder{}
	b.Grow(len(n1) + i)
	for j := 0; j < i; j++ {
		b.WriteByte('0')
	}
	b.WriteString(n1)
	return b.String()
}

func singleMultiply(num string, n byte) []int {
	result := make([]int, len(num)+1)

	b := asInt(n)

	for i := 0; i < len(num); i++ {
		a := asInt(num[i])
		value := result[i] + a*b
		result[i] = value % 10
		result[i+1] += value / 10
	}

	return result
}

func toString(result []int) string {
	b := strings.Builder{}
	b.Grow(len(result))

	skipZeros := true

	for i := len(result) - 1; i >= 0; i-- {
		if skipZeros && result[i] == 0 {
			continue
		}
		skipZeros = false
		b.WriteByte(asByte(result[i]))
	}
	return b.String()
}

func asByte(i int) byte {
	return byte(i + '0')
}

func asInt(b byte) int {
	if b < '0' || b > '9' {
		return 0
	}
	return int(b - '0')
}

func sum(items [][]int) []int {
	for i := 0; i < len(items[0]); i++ {
		for line := 1; line < len(items); line++ {
			curr := items[line]
			if i < len(curr) {
				items[0][i] += curr[i]
			}
		}
		if val := items[0][i]; val >= 10 {
			items[0][i] = val % 10
			items[0][i+1] += val / 10
		}
	}

	return items[0]
}

func reverse(s string) string {
	r := strings.Builder{}
	r.Grow(len(s))
	for i := len(s) - 1; i >= 0; i-- {
		r.WriteByte(s[i])
	}
	return r.String()
}
