package main

import "fmt"

func sortArrayByParity(A []int) []int {
	result := make([]int, len(A))
	i := 0
	j := len(A) - 1
	for k := range A {
		if A[k]%2 == 0 {
			result[i] = A[k]
			i++
		} else if A[k]%2 != 0 {
			result[j] = A[k]
			j--
		}
	}
	return result
}

func main() {
	fmt.Printf("%#v", sortArrayByParity([]int{3, 1, 2, 4}))
}
