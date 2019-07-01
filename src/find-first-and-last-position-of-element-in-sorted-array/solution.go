package main

import "fmt"

func extremeInsertionIndex(nums []int, target int, left bool) int {
	low := 0
	high := len(nums)
	for low < high {
		mid := low + (high-low)/2
		if nums[mid] > target || (left && target == nums[mid]) {
			high = mid
		} else {
			low = mid + 1
		}
	}
	return low
}

func searchRange(nums []int, target int) []int {
	targetRange := []int{-1, -1}
	leftIdx := extremeInsertionIndex(nums, target, true)

	if leftIdx == len(nums) || nums[leftIdx] != target {
		return targetRange
	}

	targetRange[0] = leftIdx
	targetRange[1] = extremeInsertionIndex(nums, target, false) - 1

	return targetRange
}

func main() {
	nums := []int{5, 7, 7, 8, 8, 10}
	fmt.Println(searchRange(nums, 7))
}
