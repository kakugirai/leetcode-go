package main

import "fmt"

func countPrimes(n int) int {
	primes := make([]bool, n)
	for i := 2; i < n; i++ {
		primes[i] = true
	}

	for i := 2; i*i < n; i++ {
		if primes[i] {
			for j := i; j*i < n; j++ {
				primes[i*j] = false
			}
		}
	}
	count := 0
	for i := range primes {
		if primes[i] {
			count++
		}
	}
	return count
}

func main() {
	fmt.Println(countPrimes(10))
}
