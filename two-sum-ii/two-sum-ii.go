package main

import "fmt"

func twoSum(numbers []int, target int) []int {
	m := make(map[int]int)
	for i := 0; i < len(numbers); i++ {
		if _, ok := m[target-numbers[i]]; ok {
			return []int{m[target-numbers[i]] + 1, i + 1}
		}
		m[numbers[i]] = i
	}
	return []int{-1, -1}
}

func main() {
	nums := []int{2, 7, 11, 15}
	target := 9
	fmt.Println(twoSum(nums, target))
}
