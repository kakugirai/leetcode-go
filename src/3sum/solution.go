package main

import (
	"fmt"
	"sort"
)

func threeSum(nums []int) [][]int {
	sort.Ints(nums)
	var result [][]int
	m := make(map[[3]int]struct{})
	for i := 1; i < len(nums)-1; i++ {
		start := 0
		end := len(nums) - 1
		for start < i && end > i {
			if nums[start]+nums[end] < 0-nums[i] {
				start++
			} else if nums[start]+nums[end] > 0-nums[i] {
				end--
			} else {
				curr := [3]int{nums[start], nums[i], nums[end]}
				if _, ok := m[curr]; !ok {
					m[curr] = struct{}{}
					result = append(result, curr[:])
				}
				start++
				end--
			}
		}
	}
	return result
}

func main() {
	//nums := []int{3, 0, -2, -1, 1, 2}
	//nums := []int{-1, 0, 1, 2, -1, 4}
	nums := []int{-4, -2, -2, -2, 0, 1, 2, 2, 2, 3, 3, 4, 4, 6, 6}
	fmt.Println(threeSum(nums))
}
