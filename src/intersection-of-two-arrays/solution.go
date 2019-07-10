package main

import "fmt"

func intersection(nums1 []int, nums2 []int) []int {
	var result []int
	m1 := make(map[int]struct{})
	m2 := make(map[int]struct{})
	for _, v := range nums1 {
		// struct{}{} doesn't occupy any additional space
		m1[v] = struct{}{}
	}
	for _, v := range nums2 {
		if _, ok := m1[v]; ok {
			m2[v] = struct{}{}
		}
	}
	for k := range m2 {
		result = append(result, k)
	}
	return result
}

func main() {
	fmt.Printf("%#v", intersection([]int{4, 5, 9}, []int{9, 4, 9, 8, 4}))
}
