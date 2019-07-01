package main

import "fmt"

// NumArray is an int array
type NumArray struct {
	sums []int
}

// Constructor initializes a numarray
func Constructor(nums []int) NumArray {
	result := NumArray{make([]int, len(nums)+1)}
	for i := range nums {
		result.sums[i+1] = result.sums[i] + nums[i]
	}
	return result
}

// SumRange returns sum from na[i] to na[j]
func (na *NumArray) SumRange(i int, j int) int {
	return na.sums[j+1] - na.sums[i]
}

/**
 * Your NumArray object will be instantiated and called as such:
 * obj := Constructor(nums);
 * param_1 := obj.SumRange(i,j);
 */

func main() {
	nums := []int{-2, 0, 3, -5, 2, -1}
	na := Constructor(nums)
	//for i := range na.sums {
	//	fmt.Println(na.sums[i])
	//}
	fmt.Println(na.SumRange(2, 5))
}
