package main

import "fmt"

func main() {
	haystack := "hello"
	needle := "l"
	fmt.Println(strStr(haystack, needle))
}
func strStr(haystack string, needle string) int {
	if needle == "" || len(needle) == 0 {
		return 0
	}

	if haystack == "" || len(needle) > len(haystack) {
		return -1
	}

	var i int
	for i = 0; i <= len(haystack)-len(needle); i++ {
		var j int
		for j = 0; j < len(needle); j++ {
			if needle[j] != haystack[i+j] {
				break
			}
		}
		if j == len(needle) {
			return i
		}
	}
	return -1
}
