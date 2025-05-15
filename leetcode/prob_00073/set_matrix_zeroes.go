package prob_00073

func setZeroes(matrix [][]int) {
	rowZeros, colZeros := zerosAt(matrix)
	for row := range rowZeros {
		clearRow(matrix, row)
	}
	for col := range colZeros {
		clearCol(matrix, col)
	}
}

func clearCol(matrix [][]int, col int) {
	for i := range matrix {
		matrix[i][col] = 0
	}
}

func clearRow(matrix [][]int, row int) {
	for i := range matrix[row] {
		matrix[row][i] = 0
	}
}

func zerosAt(matrix [][]int) (map[int]bool, map[int]bool) {
	var rowZeros, colZeros = make(map[int]bool), make(map[int]bool)
	for i := 0; i < len(matrix); i++ {
		for j := 0; j < len(matrix[i]); j++ {
			if matrix[i][j] == 0 {
				rowZeros[i] = true
				colZeros[j] = true
			}
		}
	}
	return rowZeros, colZeros
}
