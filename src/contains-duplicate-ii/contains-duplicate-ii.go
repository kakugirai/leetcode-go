package main

import "fmt"

func containsNearbyDuplicate(nums []int, k int) bool {
	m := make(map[int]int)
	for i := range nums {
		if _, ok := m[nums[i]]; ok && i-m[nums[i]] <= k {
			return true
		}
		m[nums[i]] = i
	}
	return false
}

func main() {
	nums := []int{1, 2, 3, 1, 2, 3}
	k := 2
	fmt.Println(containsNearbyDuplicate(nums, k))
}
