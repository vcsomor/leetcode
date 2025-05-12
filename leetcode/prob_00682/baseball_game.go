package prob_00682

import "strconv"

func calPoints(operations []string) int {
	records := make([]int, len(operations)*2)
	ptr := 0
	for _, op := range operations {
		if op == "+" {
			records[ptr] = add(records, ptr-2, ptr-1)
			ptr++
		} else if op == "D" {
			records[ptr] = addDouble(records, ptr-1)
			ptr++
		} else if op == "C" {
			ptr--
		} else {
			i, _ := strconv.Atoi(op)
			records[ptr] = i
			ptr++
		}
	}

	return sumRecords(records, ptr)
}

func addDouble(records []int, n int) int {
	if n >= 0 {
		return records[n] * 2
	}
	return 0
}

func add(records []int, nPrev int, n int) int {
	sum := 0
	if nPrev >= 0 {
		sum += records[nPrev]
	}
	if n >= 0 {
		sum += records[n]
	}
	return sum
}

func sumRecords(records []int, ptr int) int {
	sum := 0
	for i := 0; i < ptr; i++ {
		sum += records[i]
	}
	return sum
}
