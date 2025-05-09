// Copyright (c) 2017-2025, Cloudera, Inc. All rights reserved.
//
// Except as expressly permitted in a written agreement between you or your company
// and Cloudera, Inc. or an authorized affiliate or partner thereof, any use,
// reproduction, modification, redistribution, sharing, lending or other exploitation
// of all or any part of the contents of this software is strictly prohibited.

package prob_00459

import "strings"

func repeatedSubstringPattern(s string) bool {
	return repeatedSubstringPatternBest(s)
}

func repeatedSubstringPatternBest(s string) bool {
	if len(s) == 1 {
		return false
	}

	ss := s + s
	ss = ss[1 : len(ss)-1]
	return strings.Contains(ss, s)
}

func repeatedSubstringPatternTrivial(s string) bool {
	l := len(s)
	if l == 1 {
		return false
	}

	maxWordLength := l / 2
	for currWordLength := 1; currWordLength <= maxWordLength; currWordLength++ {
		pattern := s[:currWordLength]

		allMatch := true

		for matchRound := 1; matchRound*currWordLength < l; matchRound++ {
			at := matchRound * currWordLength
			if pattern != s[at:min(at+currWordLength, l)] {
				allMatch = false
				break
			}
		}

		if allMatch {
			return true
		}
	}

	return false
}
