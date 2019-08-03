package main

import "fmt"

func searchInsert(nums []int, target int) int {
	if len(nums) == 0 {
		return -1
	}

	if target < nums[0] {
		return 0
	}

	start := 0
	end := len(nums) - 1
	var mid int

	for start+1 < end {
		mid = start + (end-start)/2
		if nums[mid] == target {
			return mid
		} else if nums[mid] < target {
			start = mid
		} else if nums[mid] > target {
			end = mid
		}
	}

	if nums[start] == target {
		return start
	}

	if nums[end] == target {
		return end
	}

	if nums[end] < target {
		return end + 1
	}

	return start + 1
}

func main() {
	nums := []int{1, 3, 5, 6}
	target := 5
	fmt.Println(searchInsert(nums, target))
}
