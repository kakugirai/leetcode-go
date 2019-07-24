package main

import "fmt"

func permute(nums []int) [][]int {
	var result [][]int
	var curr []int
	m := make(map[int]struct{})
	generatePermutation(nums, curr, m, &result)
	return result
}

func generatePermutation(nums []int, curr []int, m map[int]struct{}, result *[][]int) {
	cpy := make([]int, len(curr))
	copy(cpy, curr)
	if len(cpy) == len(nums) {
		*result = append(*result, cpy)
	} else {
		for i := range nums {
			if _, ok := m[nums[i]]; ok {
				continue
			} else {
				m[nums[i]] = struct{}{}
				curr = append(curr, nums[i])
				generatePermutation(nums, curr, m, result)
				delete(m, curr[len(curr)-1])
				curr = curr[:len(curr)-1]
			}
		}
	}
}

func main() {
	fmt.Printf("%#v", permute([]int{5, 4, 2, 6}))
}
