package prob_01232

func checkStraightLine(coordinates [][]int) bool {
	const X, Y = 0, 1

	if len(coordinates) == 2 {
		return true
	}

	x0, x1 := coordinates[0][X], coordinates[1][X]
	y0, y1 := coordinates[0][Y], coordinates[1][Y]

	for i := 2; i < len(coordinates); i++ {
		x, y := coordinates[i][X], coordinates[i][Y]
		if (x-x0)*(y1-y0) != (y-y0)*(x1-x0) {
			return false
		}
	}
	return true
}
