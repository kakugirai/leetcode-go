package main

import (
	"fmt"
	"math"
)

func setZeroes(matrix [][]int) {
	for i := range matrix {
		for j := range matrix[i] {
			// find the zero value
			if matrix[i][j] == 0 {
				// set all non zero values in the zero value row to math.MinInt64
				for k := range matrix {
					if matrix[k][j] != 0 {
						matrix[k][j] = math.MinInt64
					}
				}
				// set all non zero values in the zero value col to math.MinInt64
				for k := range matrix[i] {
					if matrix[i][k] != 0 {
						matrix[i][k] = math.MinInt64
					}
				}
			}
		}
	}

	// find all values equals to math.MinInt64 and set them to zero
	for i := range matrix {
		for j := range matrix[i] {
			if matrix[i][j] == math.MinInt64 {
				matrix[i][j] = 0
			}
		}
	}
}

func main() {
	matrix := [][]int{
		{0, 1, 2, 0},
		{3, 4, 5, 2},
		{1, 3, 1, 5},
	}
	setZeroes(matrix)
	fmt.Println(math.MinInt64)
	fmt.Printf("%#v", matrix)
}
