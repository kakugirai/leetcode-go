package main

import "fmt"

func searchMatrix(matrix [][]int, target int) bool {
	if len(matrix) == 0 || len(matrix[0]) == 0 {
		return false
	}

	rows := len(matrix)
	cols := len(matrix[0])
	left, right := 0, rows*cols-1

	for left+1 < right {
		mid := (left + right) / 2
		midRow, midCol := mid/cols, mid%cols
		if matrix[midRow][midCol] < target {
			left = mid
		} else {
			right = mid
		}
	}

	midRow, midCol := left/cols, left%cols
	if matrix[midRow][midCol] == target {
		return true
	}

	midRow, midCol = right/cols, right%cols
	if matrix[midRow][midCol] == target {
		return true
	}

	return false
}

func main() {
	matrix := [][]int{
		{1, 3, 5, 7},
		{10, 11, 16, 20},
		{23, 30, 34, 50},
	}
	fmt.Println(searchMatrix(matrix, 3))
}
