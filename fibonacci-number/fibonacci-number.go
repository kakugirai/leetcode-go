package main

import "fmt"

var m map[int]int

func fib(n int) int {
	m = make(map[int]int, n)
	m[0], m[1] = 0, 1
	if _, ok := m[n]; ok {
		return m[n]
	}
	result := fib(n-1) + fib(n-2)
	m[n] = result
	return result
}

func main() {
	fmt.Println(fib(6))
}
