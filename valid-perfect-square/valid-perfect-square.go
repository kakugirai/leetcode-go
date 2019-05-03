package main

import "fmt"

func isPerfectSquare(num int) bool {
	if num == 1 {
		return true
	}

	left := 0
	right := num - 1

	for left < right {
		mid := left + (right-left)/2
		if mid*mid == num {
			return true
		}
		if mid*mid < num {
			left = mid + 1
		} else {
			right = mid
		}
	}
	return false
}

func main() {
	fmt.Println(isPerfectSquare(16))
}
