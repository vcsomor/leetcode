package prob_00054

var (
	moveRight = [2]int{0, 1}
	moveDown  = [2]int{1, 0}
	moveLeft  = [2]int{0, -1}
	moveUp    = [2]int{-1, 0}
	moves     = [][2]int{moveRight, moveDown, moveLeft, moveUp}

	RIGHT = 0
	DOWN  = 1
	LEFT  = 2
	UP    = 3
)

func spiralOrder(matrix [][]int) []int {
	h := len(matrix)
	w := len(matrix[0])
	n := w * h

	ret := make([]int, n)

	currentDir := RIGHT
	currentIter := w

	x, y := 0, -1

	for i := 0; i < n; i++ {
		x += moves[currentDir][0]
		y += moves[currentDir][1]
		ret[i] = matrix[x][y]

		switch currentDir {
		case RIGHT: // right
			currentIter--
			if currentIter == 0 {
				currentDir = DOWN
				currentIter = h - 1
				h--
			}

		case DOWN: // down
			currentIter--
			if currentIter == 0 {
				currentDir = LEFT
				currentIter = w - 1
				w--
			}

		case LEFT: // left
			currentIter--
			if currentIter == 0 {
				currentDir = UP
				currentIter = h - 1
				h--
			}

		case UP: // up
			currentIter--
			if currentIter == 0 {
				currentDir = RIGHT
				currentIter = w - 1
				w--
			}
		}
	}

	return ret
}
