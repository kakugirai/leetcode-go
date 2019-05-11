package main

import (
	"fmt"
)

func missingNumber(nums []int) int {
	sum := 0
	for i := range nums {
		sum += nums[i]
	}
	return len(nums)*(len(nums)+1)/2 - sum
}

func main() {
	nums := []int{9, 6, 4, 2, 3, 5, 7, 0, 1}
	fmt.Println(missingNumber(nums))
}
