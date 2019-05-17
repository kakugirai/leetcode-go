package main

import "fmt"

func validPalindrome(s string) bool {
	i := 0
	j := len(s) - 1

	for i < j {
		if s[i] != s[j] {
			return isPalindromeRange(s, i+1, j) || isPalindromeRange(s, i, j-1)
		}
		i++
		j--
	}
	return true
}

func isPalindromeRange(s string, i int, j int) bool {
	for i < j {
		if s[i] != s[j] {
			return false
		}
		i++
		j--
	}
	return true
}

func main() {
	s := "abcdcbaz"
	fmt.Println(validPalindrome(s))
}
