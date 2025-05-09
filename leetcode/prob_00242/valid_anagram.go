// Copyright (c) 2017-2025, Cloudera, Inc. All rights reserved.
//
// Except as expressly permitted in a written agreement between you or your company
// and Cloudera, Inc. or an authorized affiliate or partner thereof, any use,
// reproduction, modification, redistribution, sharing, lending or other exploitation
// of all or any part of the contents of this software is strictly prohibited.

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
