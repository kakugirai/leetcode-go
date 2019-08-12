package main

import "fmt"

func validMountainArray(A []int) bool {
	if len(A) < 3 || A[1] < A[0] {
		return false
	}
	peak := -1
	for i := 1; i < len(A); i++ {
		if A[i] < A[i-1] {
			peak = i - 1
			break
		}
	}
	if peak == -1 {
		return false
	}
	for i := peak + 1; i < len(A); i++ {
		if A[i] >= A[i-1] {
			return false
		}
	}
	return true
}

func main() {
	//A := []int{2, 1}
	A := []int{14, 82, 89, 84, 79, 70, 70, 68, 67, 66, 63, 60, 58, 54, 44, 43, 32, 28, 26, 25, 22, 15, 13, 12, 10, 8, 7, 5, 4, 3}
	//A := []int{1, 2, 3, 4, 3, 2, 1}
	fmt.Println(validMountainArray(A))
}
