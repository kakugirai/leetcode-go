package main

import "fmt"

func solveNQueens(n int) [][]string {
	var result [][]string
	queens := make([][]byte, n)
	for i := range queens {
		queens[i] = make([]byte, n)
		for j := range queens[i] {
			queens[i][j] = '.'
		}
	}
	solver(&queens, 0, n, &result)
	return result
}

func solver(queens *[][]byte, row int, n int, result *[][]string) {
	curr := make([]string, n)
	for i := range curr {
		curr[i] = string((*queens)[i])
	}

	if row == n {
		*result = append(*result, curr)
		return
	}
	for col := 0; col < n; col++ {
		if isValid(queens, row, col, n) {
			(*queens)[row][col] = 'Q'
			solver(queens, row+1, n, result)
			(*queens)[row][col] = '.'
		}
	}
}

func isValid(queens *[][]byte, row int, col int, n int) bool {
	for i := 0; i < row; i++ {
		if (*queens)[i][col] == 'Q' {
			return false
		}
	}
	for i, j := row-1, col-1; i >= 0 && j >= 0; i, j = i-1, j-1 {
		if (*queens)[i][j] == 'Q' {
			return false
		}
	}

	for i, j := row-1, col+1; i >= 0 && j < n; i, j = i-1, j+1 {
		if (*queens)[i][j] == 'Q' {
			return false
		}
	}
	return true
}

func main() {
	fmt.Printf("%#v", solveNQueens(4))
}
