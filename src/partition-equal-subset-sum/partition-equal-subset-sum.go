package main

import (
	"fmt"
)

func canPartition(nums []int) bool {
	sum := 0
	for _, v := range nums {
		sum += v
	}
	if sum&1 == 1 {
		return false
	}

	half := sum / 2
	dp := make([]bool, half+1)
	dp[0] = true

	for _, v := range nums {
		for j := half; j >= v; j-- {
			dp[j] = dp[j] || dp[j-v]
		}
	}
	return dp[half]
}

func main() {
	nums := []int{1, 5, 11, 5}
	fmt.Println(canPartition(nums))
}
