package main

func solveSudoku(board [][]byte) {
	solve(&board)
}

func solve(board *[][]byte) bool {
	for i := range *board {
		for j := range (*board)[0] {
			if (*board)[i][j] == '.' {
				for k := 1; k < 10; k++ {
					if isValid(board, i, j, byte(k+'0')) {
						(*board)[i][j] = byte(k + '0')
						if solve(board) {
							return true
						}
						(*board)[i][j] = '.'
					}
				}
				return false
			}
		}
	}
	return true
}

func isValid(board *[][]byte, row int, col int, k byte) bool {
	for i := 0; i < 9; i++ {
		if (*board)[i][col] != '.' && (*board)[i][col] == k {
			return false
		}
		if (*board)[row][i] != '.' && (*board)[row][i] == k {
			return false
		}
		if (*board)[3*(row/3)+i/3][3*(col/3)+i%3] != '.' &&
			(*board)[3*(row/3)+i/3][3*(col/3)+i%3] == k {
			return false
		}
	}
	return true
}

func main() {
	b := [][]byte{
		{'8', '3', '.', '.', '7', '.', '.', '.', '.'},
		{'6', '.', '.', '1', '9', '5', '.', '.', '.'},
		{'.', '9', '8', '.', '.', '.', '.', '6', '.'},
		{'8', '.', '.', '.', '6', '.', '.', '.', '3'},
		{'4', '.', '.', '8', '.', '3', '.', '.', '1'},
		{'7', '.', '.', '.', '2', '.', '.', '.', '6'},
		{'.', '6', '.', '.', '.', '.', '2', '8', '.'},
		{'.', '.', '.', '4', '1', '9', '.', '.', '5'},
		{'.', '.', '.', '.', '8', '.', '.', '7', '9'},
	}
	solveSudoku(b)
}
