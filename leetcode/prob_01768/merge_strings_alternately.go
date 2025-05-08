// Copyright (c) 2017-2025, Cloudera, Inc. All rights reserved.
//
// Except as expressly permitted in a written agreement between you or your company
// and Cloudera, Inc. or an authorized affiliate or partner thereof, any use,
// reproduction, modification, redistribution, sharing, lending or other exploitation
// of all or any part of the contents of this software is strictly prohibited.

package prob_01768

import "strings"

func mergeAlternately(word1 string, word2 string) string {
	if len(word1) == 0 {
		return word2
	}
	if len(word2) == 0 {
		return word1
	}

	r1 := strings.NewReader(word1)
	r2 := strings.NewReader(word2)
	b := strings.Builder{}

	readOk := true
	read := make([]byte, 1)
	for readOk {
		readOk = false
		n, _ := r1.Read(read)
		if n == 1 {
			readOk = true
			b.Write(read)
		}
		n, _ = r2.Read(read)
		if n == 1 {
			readOk = true
			b.Write(read)
		}
	}

	return b.String()
}
