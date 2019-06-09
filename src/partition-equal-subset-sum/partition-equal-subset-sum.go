package main

import (
	"fmt"
	"strconv"
)

func canPartition(nums []int) bool {
	total := 0
	for i := range nums {
		total += nums[i]
	}

	if total%2 != 0 {
		return false
	}

	m := make(map[string]bool)
	return checkPartition(nums, 0, 0, total, &m)
}

func checkPartition(nums []int, index int, sum int, total int, m *map[string]bool) bool {
	current := strconv.Itoa(index) + strconv.Itoa(sum)

	if _, ok := (*m)[current]; ok {
		return (*m)[current]
	}
	if sum*2 == total {
		return true
	}
	if sum > total/2 || index >= len(nums) {
		return false
	}

	found := checkPartition(nums, index+1, sum, total, m) || checkPartition(nums, index+1, sum+nums[index], total, m)
	(*m)[current] = found
	return found
}

func main() {
	nums := []int{1, 5, 11, 5}
	fmt.Println(canPartition(nums))
}
