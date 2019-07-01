package main

import "fmt"

func longestPalindrome(s string) string {
	if len(s) == 0 {
		return ""
	}
	start, end := 0, 0
	for i := range s {
		len1 := expandAroundCenter(s, i, i)
		len2 := expandAroundCenter(s, i, i+1)
		var length int
		// check which one is longer
		if len1 > len2 {
			length = len1
		} else {
			length = len2
		}
		// find the start and end point
		if length > end-start {
			start = i - (length-1)/2
			end = i + length/2
		}
	}
	return s[start : end+1]
}

func expandAroundCenter(s string, left int, right int) int {
	for left >= 0 && right < len(s) && s[left] == s[right] {
		left--
		right++
	}
	return right - left - 1
}

func main() {
	s := "abbababad"
	fmt.Println(longestPalindrome(s))
}
