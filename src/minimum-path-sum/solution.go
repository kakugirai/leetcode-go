package main

import "fmt"

func minPathSum(grid [][]int) int {
	if len(grid) == 0 {
		return 0
	}

	dp := make([][]int, len(grid))
	for i := range dp {
		dp[i] = make([]int, len(grid[0]))
	}

	for i := range dp {
		for j := range dp[0] {
			dp[i][j] += grid[i][j]
			if i > 0 && j > 0 {
				if dp[i-1][j] < dp[i][j-1] {
					dp[i][j] += dp[i-1][j]
				} else {
					dp[i][j] += dp[i][j-1]
				}
			} else if i > 0 {
				dp[i][j] += dp[i-1][j]
			} else if j > 0 {
				dp[i][j] += dp[i][j-1]
			}
		}
	}
	return dp[len(dp)-1][len(dp[0])-1]
}

func main() {
	grid := [][]int{
		{1, 3, 1},
		{1, 5, 1},
		{4, 2, 1},
	}
	fmt.Println(minPathSum(grid))
}
