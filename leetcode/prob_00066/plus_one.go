package prob_00066

func plusOne(digits []int) []int {
	lastIdx := len(digits) - 1

	carryOver := 0
	for i := lastIdx; i >= 0; i-- {
		currDigit := digits[i]
		if currDigit == 9 {
			carryOver = 1
			digits[i] = 0
			continue
		}
		carryOver = 0
		digits[i] = currDigit + 1
		break
	}

	if carryOver == 0 {
		return digits
	}

	return make10Exponent(len(digits) + 1)
}

func make10Exponent(len int) []int {
	num := make([]int, len)
	num[0] = 1
	return num
}
