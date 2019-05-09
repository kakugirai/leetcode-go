package main

import "fmt"

func climbStairs(n int) int {
	m := make(map[int]int)
	m[0], m[1] = 1, 1
	for i := 2; i <= n; i++ {
		m[i] = m[i-1] + m[i-2]
	}
	return m[n]
}

func main() {
	fmt.Println(climbStairs(5))
}
