package main

import "fmt"

func main() {
	nums := []int{1, 2, 3, 4, 5, 6, 7}
	target := 3
	result := twoSum(nums, target)
	for _, v := range result {
		fmt.Printf("%d", v)
	}
}

func twoSum(nums []int, target int) []int {
	record := make(map[int]int)
	for i, j := range nums {
		complement := target - j
		res, ok := record[complement]
		if ok {
			return []int{res, i}
		}
		record[j] = i
	}
	return []int{}
}

