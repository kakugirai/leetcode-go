package main

import "fmt"

func nextPermutation(nums []int) {
	if len(nums) == 0 || len(nums) == 1 {
		return
	}
	low := -1
	for i := 1; i < len(nums); i++ {
		if nums[i] > nums[i-1] {
			low = i - 1
		}
	}

	if low == -1 {
		reverse(nums, 0)
		//fmt.Println(nums)
		return
	}
	next := low + 1
	for i := low + 1; i < len(nums); i++ {
		if nums[i] <= nums[next] && nums[i] > nums[low] {
			next = i
		}
	}
	swap(nums, low, next)
	reverse(nums, low+1)
}

func reverse(nums []int, low int) {
	a := nums[low:]
	for i := len(a)/2 - 1; i >= 0; i-- {
		opp := len(a) - 1 - i
		a[i], a[opp] = a[opp], a[i]
	}
}
func swap(nums []int, i int, j int) {
	nums[i], nums[j] = nums[j], nums[i]
}

func main() {
	//nums := []int{1, 3, 2}
	nums := []int{2, 3, 1, 3, 3}
	//nums := []int{3, 2, 1}
	//nums := []int{1, 2, 5, 4, 3}
	nextPermutation(nums)
	fmt.Printf("%#v", nums)
}
