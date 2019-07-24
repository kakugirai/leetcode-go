package main

import "fmt"

func plusOne(digits []int) []int {
	if len(digits) == 0 {
		return []int{1}
	}
	digits[len(digits)-1]++

	for i := len(digits) - 1; i > 0; i-- {
		if digits[i] > 9 {
			digits[i-1]++
			digits[i] -= 10
		}
	}

	if digits[0] > 9 {
		digits[0] -= 10
		digits = append([]int{1}, digits...)
	}
	return digits
}

func main() {
	digits := []int{7, 2, 8, 5, 0, 9, 1, 2, 9, 5, 3, 6, 6, 7, 3, 2, 8, 4, 3, 7, 9, 5, 7, 7, 4, 7, 4, 9, 4, 7, 0, 1, 1, 1, 7, 4, 0, 0, 6}
	fmt.Printf("%#v", plusOne(digits))
}
