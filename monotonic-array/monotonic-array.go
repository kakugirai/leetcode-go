package main

import "fmt"

func isMonotonic(A []int) bool {
	if len(A) == 0 || len(A) == 1 {
		return true
	}
	if A[0] <= A[len(A)-1] {
		for i := 1; i < len(A); i++ {
			if A[i] < A[i-1] {
				return false
			}
		}
	} else {
		for i := 1; i < len(A); i++ {
			if A[i] > A[i-1] {
				return false
			}
		}
	}
	return true
}

func main() {
	A := []int{1, 2, 2, 5}
	fmt.Println(isMonotonic(A))
}
