package main

import "fmt"

func isHappy(n int) bool {
	seen := make(map[int]struct{})
	for n != 1 {
		sum := 0
		for n > 0 {
			sum += (n % 10) * (n % 10)
			n /= 10
		}

		if _, ok := seen[sum]; ok {
			return false
		}
		seen[sum] = struct{}{}
		n = sum
	}
	return true
}

func main() {
	fmt.Println(isHappy(19))
}
