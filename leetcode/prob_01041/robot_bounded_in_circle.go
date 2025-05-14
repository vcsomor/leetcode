package prob_01041

const DIRECTIONS = 4

var (
	moveU = []int{0, -1}
	moveL = []int{-1, 0}
	moveR = []int{1, 0}
	moveD = []int{0, 1}

	rotation = [DIRECTIONS][]int{
		moveU, // up
		moveR, // right
		moveD, // down
		moveL, // left
	}
)

func isRobotBounded(instructions string) bool {
	// 0 = up, 1 = left, 2 = down, 3 = right
	currentDir := 0

	// x, y coordinates
	x := 0
	y := 0

	for _, i := range instructions {
		switch i {
		case 'G':
			x += rotation[currentDir][0]
			y += rotation[currentDir][1]
			continue
		case 'L':
			currentDir = (currentDir + 3) % DIRECTIONS
			continue
		case 'R':
			currentDir = (currentDir + 1) % DIRECTIONS
			continue
		}
	}

	// after processing the instruction there are two cases
	// 1. if the robot is at the origin (0,0) then it is bounded
	if x == 0 && y == 0 {
		return true
	}

	// 2. if the robot is not at the origin, then it should NOT be facing up
	return currentDir != 0
}
