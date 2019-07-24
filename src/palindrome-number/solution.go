package main

import "fmt"

func isPalindrome(x int) bool {
	if x < 0 || (x%10 == 0 && x != 0) {
		return false
	}
	r := 0

	// even digit case: x = 1221, r = 0;, x = 122, r = 1; x = 12, r = 12
	// odd digit case: x = 12321, r = 0;, x = 1232, r = 1; x = 123, r = 12; x = 12, r = 123
	for x > r {
		r = r*10 + x%10
		x /= 10
	}
	return x == r || x == r/10
}

func main() {
	fmt.Println(isPalindrome(11011))
}
