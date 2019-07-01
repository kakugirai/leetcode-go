package main

import (
	"fmt"
)

func subsets(nums []int) [][]int {
	var result [][]int
	var curr []int
	generateSubsets(nums, 0, &curr, &result)
	return result
}

func generateSubsets(nums []int, idx int, curr *[]int, result *[][]int) {
	cpy := make([]int, len(*curr))
	copy(cpy, *curr)
	*result = append(*result, cpy)
	for i := idx; i < len(nums); i++ {
		*curr = append(*curr, nums[i])
		generateSubsets(nums, i+1, curr, result)
		*curr = (*curr)[:len(*curr)-1]
	}
}

func main() {
	nums := []int{1, 2, 3}
	fmt.Println(subsets(nums))
}
