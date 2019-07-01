package main

import "fmt"

func singleNumber(nums []int) int {
	sum := 0
	singleSum := 0
	m := make(map[int]struct{})
	for _, v := range nums {
		sum += v
		if _, ok := m[v]; !ok {
			singleSum += v
			m[v] = struct{}{}
		}
	}
	return singleSum*2 - sum
}

func main() {
	nums := []int{4, 1, 2, 1, 2}
	fmt.Println(singleNumber(nums))
}
