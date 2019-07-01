package main

import "fmt"

func minDistance(word1 string, word2 string) int {
	if word1 == word2 {
		return 0
	}
	if len(word1) == 0 || len(word2) == 0 {
		if len(word1)-len(word2) < 0 {
			return len(word2) - len(word1)
		}
		return len(word1) - len(word2)
	}

	matrix := make([][]int, len(word2)+1)
	for i := range matrix {
		matrix[i] = make([]int, len(word1)+1)
	}

	for i := range word2 {
		matrix[i+1][0] = i + 1
	}
	for i := range word1 {
		matrix[0][i+1] = i + 1
	}

	for i := range word2 {
		for j := range word1 {
			if word2[i] == word1[j] {
				matrix[i+1][j+1] = matrix[i][j]
			} else {
				arr := []int{matrix[i][j], matrix[i+1][j], matrix[i][j+1]}
				matrix[i+1][j+1] = getMin(arr) + 1
			}
		}
	}
	return matrix[len(word2)][len(word1)]
}

func getMin(arr []int) int {
	min := 2<<31 - 1
	for i := range arr {
		if arr[i] < min {
			min = arr[i]
		}
	}
	return min
}

func main() {
	word1 := "horse"
	word2 := "ros"
	fmt.Println(minDistance(word1, word2))
}
