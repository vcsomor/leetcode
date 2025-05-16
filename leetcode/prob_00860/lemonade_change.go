package prob_00860

func lemonadeChange(bills []int) bool {
	// only interested in fives and tens
	fives, tens := 0, 0

	for _, bill := range bills {
		switch bill {
		case 5:
			fives++

		case 10:
			if fives >= 1 {
				tens++
				fives--
				continue
			}
			return false

		case 20:
			// proiritise using 10 and 5
			if tens >= 1 && fives >= 1 {
				tens--
				fives--
				continue
			}

			// fallback to using 5s
			if fives >= 3 {
				fives -= 3
				continue
			}

			// No change
			return false
		}
	}

	return true
}
