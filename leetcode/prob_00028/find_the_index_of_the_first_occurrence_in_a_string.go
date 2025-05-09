// Copyright (c) 2017-2025, Cloudera, Inc. All rights reserved.
//
// Except as expressly permitted in a written agreement between you or your company
// and Cloudera, Inc. or an authorized affiliate or partner thereof, any use,
// reproduction, modification, redistribution, sharing, lending or other exploitation
// of all or any part of the contents of this software is strictly prohibited.

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
