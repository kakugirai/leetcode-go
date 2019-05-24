package main

import "fmt"

//https://leetcode.com/problems/search-in-rotated-sorted-array/discuss/154836/The-INF-and-INF-method-but-with-a-better-explanation-for-dummies-like-me
func search(nums []int, target int) int {
	if len(nums) == 0 || nums == nil {
		return -1
	}
	left := 0
	right := len(nums) - 1

	for left <= right {
		mid := (left + right) >> 1
		comparator := nums[mid]
		if (target < nums[0]) && (nums[mid] < nums[0]) || (target >= nums[0]) && (nums[mid] >= nums[0]) {
			comparator = nums[mid]
		} else {
			if target < nums[0] {
				comparator = -int(^uint(0)>>1) - 1
			} else {
				comparator = int(^uint(0) >> 1)
			}
		}

		if target == comparator {
			return mid
		}
		if target > comparator {
			left = mid + 1
		} else if target < comparator {
			right = mid - 1
		}
	}
	return -1
}

func main() {
	nums := []int{4, 5, 6, 7, 0, 1, 2}
	target := 1
	fmt.Println(search(nums, target))
}
