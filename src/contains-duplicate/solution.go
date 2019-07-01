package main

import "fmt"

func containsDuplicate(nums []int) bool {
	m := make(map[int]struct{})
	for _, v := range nums {
		if _, ok := m[v]; ok {
			return true
		}
		m[v] = struct{}{}
	}
	return false
}

func main() {
	//nums := []int{1, 1, 1, 3, 3, 4, 3, 2, 4, 2}
	nums := []int{1, 2, 3, 4}
	fmt.Println(containsDuplicate(nums))
}
