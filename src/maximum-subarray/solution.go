package main

import (
	"fmt"
)

func main() {
	nums := []int{-2, 1, -3, 4, -1, 2, 1, -5, 4}
	fmt.Println(maxSubArray(nums))
}

func maxSubArray(nums []int) int {
	// Time complexity: O(n)
	// Space complexity: O(2)
	// Apparently we can reduce space complexity to O(1) if we choose an in-place solution.
	// But I found this more readable for me so why don't we just consume more memory :)
	localMax := nums[0]
	globalMax := nums[0]
	for i := 1; i < len(nums); i++ {
		if localMax <= 0 {
			localMax = nums[i]
		} else {
			localMax += nums[i]
		}
		if localMax > globalMax {
			globalMax = localMax
		}
	}
	return globalMax
}
