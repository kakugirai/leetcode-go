package main

import (
	"fmt"
)

func reverse(x int) int {
	negative := false
	if x < 0 {
		negative = true
		x *= -1
	}
	reversed := 0
	for x > 0 {
		reversed = reversed*10 + x%10
		x /= 10
	}

	if negative {
		reversed *= -1
	}

	if reversed >= (2<<31-1) || reversed <= -1*(2<<31) {
		return 0
	} else {
		return reversed
	}
}

func main() {
	result := reverse(-210)
	fmt.Println(reverse(result))
	fmt.Println(2<<31 - 1)
	fmt.Println(2147483651)
}
