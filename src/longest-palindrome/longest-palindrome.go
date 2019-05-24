package main

import "fmt"

func longestPalindrome(s string) int {
	var result int
	m := make(map[byte]int)
	for i := range s {
		m[s[i]]++
	}

	haveOddNumber := false
	for _, v := range m {
		if v%2 == 0 {
			result += v
		}
		if v%2 == 1 {
			haveOddNumber = true
			result += (v / 2) * 2
		}
	}
	if haveOddNumber {
		return result + 1
	}
	return result
}

func main() {
	s := "abbbccccdd"
	fmt.Println(longestPalindrome(s))
}
