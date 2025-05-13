package prob_template

func judgeCircle(moves string) bool {
	var ups, downs, lefts, rights int
	for _, m := range moves {
		switch m {
		case 'U':
			ups++
		case 'D':
			downs++
		case 'L':
			lefts++
		case 'R':
			rights++
		}
	}
	return (ups == downs) && (lefts == rights)
}
