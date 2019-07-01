package main

import "fmt"

func majorityElement(nums []int) int {
	if len(nums) == 1 {
		return nums[0]
	}
	m := make(map[int]int)
	for i := range nums {
		if _, ok := m[nums[i]]; ok {
			m[nums[i]]++
			if m[nums[i]] > len(nums)/2 {
				return nums[i]
			}
		} else {
			m[nums[i]] = 1
		}
	}
	return 0
}

func main() {
	nums := []int{2, 2, 1, 1, 1, 2, 2}
	fmt.Println(majorityElement(nums))
}
