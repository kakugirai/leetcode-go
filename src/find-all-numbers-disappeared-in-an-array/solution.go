package main

import (
	"fmt"
	"math"
)

func findDisappearedNumbers(nums []int) []int {
	var missing []int

	for i := range nums {
		val := int(math.Abs(float64(nums[i])) - 1)
		if nums[val] > 0 {
			nums[val] = -nums[val]
		}
	}

	for i := range nums {
		if nums[i] > 0 {
			missing = append(missing, i+1)
		}
	}
	return missing
}

func main() {
	nums := []int{4, 3, 2, 7, 8, 2, 3, 1}
	fmt.Println(findDisappearedNumbers(nums))
}
