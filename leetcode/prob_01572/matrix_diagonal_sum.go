package prob_01572

func diagonalSum(mat [][]int) int {
	dim := len(mat)
	maxIdx := dim - 1
	sum := 0

	for i := 0; i < dim; i++ {
		sum += mat[i][i]        // primary diagonal
		sum += mat[i][maxIdx-i] // secondary diagonal
	}

	if dim&1 == 1 { // if odd
		sum -= mat[dim/2][dim/2]
	}

	return sum
}
