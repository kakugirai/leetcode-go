package main

import "fmt"

func helper(n int, m map[int]int) int {
	m = make(map[int]int, n)
	m[0], m[1] = 0, 1
	if _, ok := m[n]; ok {
		return m[n]
	}
	result := helper(n-1, m) + helper(n-2, m)
	m[n] = result
	return result
}

func fib(n int) int {
	return helper(n, map[int]int{})
}

func main() {
	fmt.Println(fib(6))
}
