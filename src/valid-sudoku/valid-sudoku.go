package main

import "fmt"

func isValidSudoku(board [][]byte) bool {
	for row := 0; row < 9; row++ {
		if !isValidRow(board, row) {
			return false
		}
	}

	for col := 0; col < 9; col++ {
		if !isValidCol(board, col) {
			return false
		}
	}

	for pod := 0; pod < 9; pod++ {
		if !isValidPod(board, pod) {
			return false
		}
	}
	return true
}

func isValidRow(board [][]byte, row int) bool {
	var nums [10]bool
	for col := 0; col < 9; col++ {
		n := convertToNumber(board[row][col])
		if n < 0 {
			continue
		}
		if nums[n] {
			return false
		}
		nums[n] = true
	}
	return true
}

func isValidCol(board [][]byte, col int) bool {
	var nums [10]bool
	for row := 0; row < 9; row++ {
		n := convertToNumber(board[row][col])
		if n < 0 {
			continue
		}
		if nums[n] {
			return false
		}
		nums[n] = true
	}
	return true
}

func isValidPod(board [][]byte, pod int) bool {
	var nums [10]bool

	row := (pod / 3) * 3
	col := (pod % 3) * 3

	for i := 0; i < 3; i++ {
		for j := 0; j < 3; j++ {
			n := convertToNumber(board[row+i][col+j])
			if n < 0 {
				continue
			}
			if nums[n] {
				return false
			}
			nums[n] = true
		}
	}
	return true
}

func convertToNumber(b byte) int {
	if b == '.' {
		return -1
	}
	return int(b - '0')
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
	fmt.Println(isValidSudoku(b))
}
