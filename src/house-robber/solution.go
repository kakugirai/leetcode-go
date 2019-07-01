package main

import "fmt"

func rob(nums []int) int {
	if nums == nil || len(nums) == 0 {
		return 0
	}
	if len(nums) == 1 {
		return nums[0]
	}
	if len(nums) == 2 {
		if nums[0] > nums[1] {
			return nums[0]
		}
		return nums[1]
	}

	dp := make([]int, len(nums))
	dp[0] = nums[0]
	if nums[0] > nums[1] {
		dp[1] = nums[0]
	} else {
		dp[1] = nums[1]
	}

	for i := 2; i < len(dp); i++ {
		if nums[i]+dp[i-2] > dp[i-1] {
			dp[i] = nums[i] + dp[i-2]
		} else {
			dp[i] = dp[i-1]
		}
	}
	return dp[len(nums)-1]
}

func main() {
	nums := []int{1, 2, 3, 1}
	fmt.Println(rob(nums))
}
