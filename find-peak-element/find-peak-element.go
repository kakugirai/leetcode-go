package main

import "fmt"

func findPeakElement(nums []int) int {
	left := 0
	right := len(nums) - 1

	for left < right {
		mid := left + (right-left)/2
		if nums[mid] < nums[mid+1] {
			left = mid + 1
		} else {
			right = mid
		}
	}
	return left
}

func main() {
	fmt.Println(findPeakElement([]int{1, 3, 4, 5, 6, 4, 6}))
}
