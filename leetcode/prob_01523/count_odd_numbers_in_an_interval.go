package prob_01523

func countOdds(low int, high int) int {
	cnt := (high - low) >> 1
	if (low&1 == 1) || (high&1 == 1) {
		cnt++
	}
	return cnt
}
