package main

import (
	"fmt"
	"math"
)

func isSquare(n int) bool {
	nsqrt := int(math.Sqrt(float64(n)))
	return nsqrt*nsqrt == n
}

func numSquares(n int) int {
	// Based on Lagrange's Four Square theorem, there
	// are only 4 possible results: 1, 2, 3, 4.
	if isSquare(n) {
		return 1
	}
	for (n & 3) == 0 {
		n >>= 2
	}

	// The result is 4 if and only if n can be written in the
	// form of 4^k*(8*m + 7). Please refer to
	// Legendre's three-square theorem.
	if (n & 7) == 7 {
		return 4
	}

	// Check whether 2 is the result.
	nsqrt := int(math.Sqrt(float64(n)))
	for i := 1; i <= nsqrt; i++ {
		if isSquare(n - i*i) {
			return 2
		}
	}
	return 3
}

func main() {
	fmt.Println(numSquares(13))
}
