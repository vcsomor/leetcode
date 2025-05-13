package prob_01275

var seq = []byte{'X', 'O', 'X', 'O', 'X', 'O', 'X', 'O', 'X'}

var winningPatterns = [][]int{
	{0, 1, 2},
	{3, 4, 5},
	{6, 7, 8},

	{0, 3, 6},
	{1, 4, 7},
	{2, 5, 8},

	{0, 4, 8},
	{2, 4, 6},
}

func tictactoe(moves [][]int) string {
	flattened := transformMoves(moves)

	if hasWon(flattened, 'X') {
		return "A"
	}

	if hasWon(flattened, 'O') {
		return "B"
	}

	if isPending(flattened) {
		return "Pending"
	}

	return "Draw"
}

func isPending(flattened []byte) bool {
	for _, b := range flattened {
		if b == 0 {
			return true
		}
	}
	return false
}

func hasWon(flattened []byte, who byte) bool {
	for _, wp := range winningPatterns {
		p0 := flattened[wp[0]]
		p1 := flattened[wp[1]]
		p2 := flattened[wp[2]]
		if p0 == who && p1 == who && p2 == who {
			return true
		}
	}
	return false
}

func transformMoves(moves [][]int) []byte {
	m := make([]byte, 9)
	for i := range moves {
		row := moves[i][0]
		col := moves[i][1]
		m[row*3+col] = seq[i]
	}
	return m
}
