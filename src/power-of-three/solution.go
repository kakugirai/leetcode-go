package main

import "fmt"

func isPowerOfThree(n int) bool {
	return n > 0 && 1162261467%n == 0
}

func main() {
	fmt.Println(isPowerOfThree(9))
}
