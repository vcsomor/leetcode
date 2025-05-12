package prob_00028

func isAnagram(s string, t string) bool {
	if len(s) != len(t) {
		return false
	}

	return matchOccurrences(countLetters(s), countLetters(t))
}

func matchOccurrences(dict0 map[byte]int, dict1 map[byte]int) bool {
	for k0, v0 := range dict0 {
		v1, ok := dict1[k0]
		if !ok {
			return false
		}
		if v0 != v1 {
			return false
		}
	}
	return true
}

func countLetters(s string) map[byte]int {
	cnt := make(map[byte]int)
	for i := range s {
		b := s[i]
		if count, ok := cnt[b]; ok {
			cnt[b] = count + 1
		} else {
			cnt[b] = 1
		}
	}
	return cnt
}
