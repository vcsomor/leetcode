package prob_00058

import "strings"

func lengthOfLastWord(s string) int {
	spl := strings.Split(s, " ")
	for i := len(spl) - 1; i >= 0; i-- {
		if len(spl[i]) == 0 {
			continue
		}
		return len(spl[i])
	}
	return 0
}
