package main

import (
	"fmt"
)

func maxProfit(prices []int) int {
	max := 0
	min := 2<<31 - 1
	for i := 0; i < len(prices); i++ {
		if prices[i] < min {
			min = prices[i]
		} else {
			if max < prices[i]-min {
				max = prices[i] - min
			}
		}
	}
	return max
}

func main() {
	prices := []int{7, 1, 5, 3, 6, 4}
	fmt.Println(maxProfit(prices))
}
