package main

import (
	"fmt"
)

func hasAlternatingBits(n int) bool {
	last := n % 2
	n >>= 1
	for n > 0 {
		current := n % 2
		if current == last {
			return false
		}
		last = current
		n >>= 1
	}
	return true
}

func main() {
	fmt.Println(hasAlternatingBits(5))
}
