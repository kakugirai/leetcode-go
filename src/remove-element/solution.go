package main

import "fmt"

func removeElement(nums []int, val int) int {
	i := 0
	for j := range nums {
		if nums[j] != val {
			nums[i] = nums[j]
			i++
		}
	}
	return i
}

func main() {
	nums := []int{3, 2, 2, 3}
	val := 3
	fmt.Println(removeElement(nums, val))
}
