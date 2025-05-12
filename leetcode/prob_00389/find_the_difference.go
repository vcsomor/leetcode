package prob_00389

func findTheDifference(s string, t string) byte {
	return findTheDifferenceXor(s, t)
}

func findTheDifferenceTrivial(s string, t string) byte {
	if len(t) == 1 {
		return t[0]
	}

	originalCnt := countLetters(s)
	addedCnt := countLetters(t)
	for letter, addedOccur := range addedCnt {
		originalOccur, origExists := originalCnt[letter]
		if !origExists || (originalOccur != addedOccur) {
			return letter
		}
	}
	return 0
}

func findTheDifferenceXor(s string, t string) byte {
	var result byte = 0

	for _, letter := range s {
		result ^= byte(letter)
	}

	for _, letter := range t {
		result ^= byte(letter)
	}

	return result
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
