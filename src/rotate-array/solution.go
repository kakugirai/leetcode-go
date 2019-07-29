package main

import "fmt"

func rotate(nums []int, k int) []int {
	if k >= len(nums) {
		return nums
	}
	nums = append(nums[k+1:], nums[0:k+1]...)
	return nums
}

func main() {
	nums := []int{1, 2, 3, 4, 5, 6, 7}
	k := 3
	fmt.Printf("%#v", rotate(nums, k))
}
