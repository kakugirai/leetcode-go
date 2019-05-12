package main

import "fmt"

func flipAndInvertImage(A [][]int) [][]int {
	for i := 0; i < len(A); i++ {
		j := 0
		k := len(A) - 1
		for j < k {
			temp := A[i][j]
			A[i][j] = A[i][k]
			A[i][k] = temp
			j++
			k--
		}

		for j = 0; j < len(A); j++ {
			A[i][j] ^= 1
		}
	}
	return A
}

func main() {
	A := [][]int{
		{1, 1, 0, 0},
		{1, 0, 0, 1},
		{0, 1, 1, 1},
		{1, 0, 1, 0},
	}

	fmt.Println(flipAndInvertImage(A))
}
