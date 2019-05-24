package main

import (
	"fmt"
	"strings"
)

func isLetterOrDigit(c byte) bool {
	if (c >= 'a' && c <= 'z') || (c >= '0' && c <= '9') {
		return true
	}
	return false
}

func isPalindrome(s string) bool {
	i := 0
	j := len(s) - 1
	s = strings.ToLower(s)

	for i < j {
		for !isLetterOrDigit(s[i]) && i < j {
			i++
		}
		for !isLetterOrDigit(s[j]) && i < j {
			j--
		}
		if s[i] != s[j] {
			return false
		}
		i++
		j--
	}
	return true
}

func main() {
	s := "A man, a plan, a canal: Panama"
	fmt.Println(isPalindrome(s))
}
