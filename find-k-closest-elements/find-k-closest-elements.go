package main

import "fmt"

func findClosestElements(arr []int, k int, x int) []int {
	left := 0
	right := len(arr) - k
	for left < right {
		mid := left + (right-left)/2
		if x-arr[mid] > arr[mid+k]-x {
			left = mid + 1
		} else {
			right = mid
		}
	}
	return arr[left : left+k]
}

func main() {
	arr := []int{1, 2, 3, 4, 5}
	fmt.Println(findClosestElements(arr, 4, 3))
}
