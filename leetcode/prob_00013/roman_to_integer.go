// Copyright (c) 2017-2025, Cloudera, Inc. All rights reserved.
//
// Except as expressly permitted in a written agreement between you or your company
// and Cloudera, Inc. or an authorized affiliate or partner thereof, any use,
// reproduction, modification, redistribution, sharing, lending or other exploitation
// of all or any part of the contents of this software is strictly prohibited.

package prob_00013

var romanPriorityMapping = map[string]int{
	"IV": 4,
	"IX": 9,
	"XL": 40,
	"XC": 90,
	"CD": 400,
	"CM": 900,
}

var romanMapping = map[string]int{
	"I": 1,
	"V": 5,
	"X": 10,
	"L": 50,
	"C": 100,
	"D": 500,
	"M": 1000,
}

func romanToInt(s string) int {
	sum := 0

	for ptr := len(s) - 1; ptr >= 0; ptr-- {
		if ptr > 0 {
			prioVal, ok := romanPriorityMapping[s[ptr-1:ptr+1]]
			if ok {
				sum += prioVal
				ptr--
				continue
			}
		}
		sum += romanMapping[s[ptr:ptr+1]]
	}
	return sum
}
