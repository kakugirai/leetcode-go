package main

import (
	"fmt"
)

func hammingDistance(x int, y int) int {
	result := 0
	for x > 0 || y > 0 {
		result += (x % 2) ^ (y % 2)
		x >>= 1
		y >>= 1
	}
	return result
}

func main() {
	fmt.Println(hammingDistance(1, 4))
}
