package prob_01672

func maximumWealth(accounts [][]int) int {
	maxWealth := 0

	for _, account := range accounts {
		maxWealth = max(maxWealth, sumOf(account))

	}

	return maxWealth
}

func sumOf(account []int) int {
	sum := 0
	for _, v := range account {
		sum += v
	}
	return sum
}
