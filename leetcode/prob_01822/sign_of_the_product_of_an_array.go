package prob_01822

func arraySign(nums []int) int {
	minusSign := false

	for _, n := range nums {
		if n == 0 {
			return 0
		}
		if n < 0 {
			minusSign = !minusSign
		}
	}

	if minusSign {
		return -1
	}
	return 1
}
