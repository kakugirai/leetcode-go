package main

import (
	"fmt"
)

func maxArea(height []int) int {
	i, j := 0, len(height)-1
	max := 0
	for i < j {
		min := 0
		if height[i] <= height[j] {
			min = height[i]
		} else {
			min = height[j]
		}
		if max < min*(j-i) {
			max = min * (j - i)
		}
		if height[i] < height[j] {
			i++
		} else {
			j--
		}
	}
	return max
}

func main() {
	height := []int{1, 1}
	fmt.Println(maxArea(height))
}
