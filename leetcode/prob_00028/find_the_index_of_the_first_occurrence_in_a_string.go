package prob_00028

func strStr(haystack string, needle string) int {
	lenHaystack := len(haystack)
	lenNeedle := len(needle)
	if lenNeedle > lenHaystack {
		return -1
	}

	for i := 0; i < (lenHaystack - lenNeedle + 1); i++ {
		if haystack[i] != needle[0] {
			continue
		}

		matchIndex := 1
		for ; matchIndex < lenNeedle; matchIndex++ {
			lowHaystack := i + matchIndex
			lowNeedle := matchIndex

			highHaystack := i + (lenNeedle - matchIndex)
			highNeedle := lenNeedle - matchIndex

			if haystack[lowHaystack] != needle[lowNeedle] {
				break
			}
			if haystack[highHaystack] != needle[highNeedle] {
				break
			}
		}
		if matchIndex == lenNeedle {
			return i
		}
	}

	return -1
}
