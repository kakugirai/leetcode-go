package main

import "fmt"

func singleNumber(nums []int) int {
	m := make(map[int]int)
	for i := range nums {
		m[nums[i]]++
	}
	for i, v := range m {
		if v == 1 {
			return i
		}
	}
	return -1
}

func main() {
	nums := []int{0, 1, 0, 1, 0, 1, 99}
	fmt.Println(singleNumber(nums))
}
