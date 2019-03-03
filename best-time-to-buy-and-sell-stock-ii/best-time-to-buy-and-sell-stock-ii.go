package main

import "fmt"

func main() {
	prices := []int{7, 1, 5, 3, 6, 4, 5}
	fmt.Println(maxProfit(prices))
}

func maxProfit(prices []int) int {
	profit := 0

	if len(prices) <= 1 {
		return profit
	}

	for i := 1; i < len(prices); i++ {
		if prices[i] > prices[i-1] {
			profit += prices[i] - prices[i-1]
		}
	}
	return profit
}
