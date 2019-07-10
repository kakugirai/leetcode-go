package main

import "fmt"

func intersect(nums1 []int, nums2 []int) []int {
	m := make(map[int]int)
	var res []int
	for _, v := range nums1 {
		m[v]++
	}
	for _, v := range nums2 {
		if m[v] > 0 {
			res = append(res, v)
			m[v]--
		}
	}
	return res
}

func main() {
	nums1 := []int{1, 2, 2, 1}
	nums2 := []int{2, 2}
	fmt.Printf("%#v", intersect(nums1, nums2))
}
