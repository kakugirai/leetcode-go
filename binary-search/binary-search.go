package main

import "fmt"

func search(nums []int, target int) int {
	left := 0
	right := len(nums) - 1
	for left <= right {
		pivot := (right + left) / 2
		if target == nums[pivot] {
			return pivot
		}
		if target < nums[pivot] {
			right = pivot - 1
		}
		if target > nums[pivot] {
			left = pivot + 1
		}
	}
	return -1
}

func main() {
	nums := []int{-1, 0, 3, 5, 9, 12}
	fmt.Println(search(nums, 9))
}
