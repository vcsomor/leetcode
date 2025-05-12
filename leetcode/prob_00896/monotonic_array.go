package prob_00869

const (
	decreasing int = -1
	stay       int = 0
	increasing int = 1
)

func isMonotonic(nums []int) bool {
	lastTendency := stay
	for i := 2; i < len(nums); i++ {
		if lastTendency == stay {
			lastTendency = tendency(nums[i-2], nums[i-1])
		}
		currTendency := tendency(nums[i-1], nums[i])
		if currTendency == stay {
			continue
		}

		if lastTendency == stay {
			lastTendency = currTendency
			continue
		}

		if currTendency != lastTendency {
			return false
		}
	}
	return true
}

func tendency(n0, n1 int) int {
	if n0 < n1 {
		return increasing
	}
	if n0 > n1 {
		return decreasing
	}
	return stay
}
