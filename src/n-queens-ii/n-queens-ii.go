package main

import "fmt"

func totalNQueens(n int) int {
	cols := make([]bool, n)
	d1 := make([]bool, n*2)
	d2 := make([]bool, n*2)
	return solver(0, cols, d1, d2, n, 0)
}

func solver(row int, cols []bool, d1 []bool, d2 []bool, n int, count int) int {
	if row == n {
		count++
	} else {
		for col := 0; col < n; col++ {
			id1 := row - col + n
			id2 := row + col
			if cols[col] || d1[id1] || d2[id2] {
				continue
			}
			cols[col] = true
			d1[id1] = true
			d2[id2] = true
			count = solver(row+1, cols, d1, d2, n, count)
			cols[col] = false
			d1[id1] = false
			d2[id2] = false
		}
	}
	return count
}

func main() {
	fmt.Println(totalNQueens(4))
}
