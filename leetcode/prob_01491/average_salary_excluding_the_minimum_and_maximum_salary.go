package prob_01491

import (
	"math"
	"sort"
)

func average(salary []int) float64 {
	l := len(salary)
	sort.Ints(salary)
	sum := 0
	for i := 1; i < l-1; i++ {
		sum += salary[i]
	}

	return round(float64(sum) / float64(l-2))
}

func round(avg float64) float64 {
	r := math.Pow10(5)
	return math.Round(avg*r) / r
}
