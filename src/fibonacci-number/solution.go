package main

import "fmt"

func helper(n int, m map[int]int) int {
	m[0], m[1] = 0, 1
	if _, ok := m[n]; ok {
		return m[n]
	}
	m[n] = helper(n-1, m) + helper(n-2, m)
	return m[n]
}

func fib(n int) int {
	m := make(map[int]int, n)
	return helper(n, m)
}

func fib2(n int) int {
	var result int
	ch, quit := make(chan int), make(chan int)
	go func(ch, quit chan int) {
		i, j := 0, 1
		for {
			select {
			case ch <- j:
				i, j = j, i+j
			case <-quit:
				return
			}
		}
	}(ch, quit)
	for i := 0; i < n; i++ {
		result = <-ch
	}
	quit <- 0
	return result
}

func main() {
	fmt.Println(fib(100))
	fmt.Println(fib2(100))
}
