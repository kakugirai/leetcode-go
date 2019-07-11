package main

import (
	"fmt"
	"strings"
)

// Approach 1: Horizontal scanning
func longestCommonPrefix(strs []string) string {
	// check if strs empty
	if len(strs) == 0 {
		return ""
	}
	// set the first str as prefix
	prefix := strs[0]
	// check if the following strs contain prefix
	for _, v := range strs[1:] {
		// trim the last char of prefix if the str does not contain prefix
		for strings.Index(v, prefix) != 0 {
			prefix = prefix[0 : len(prefix)-1]
		}
	}
	return prefix
}

// Approach 2: Vertical scanning
func longestCommonPrefix2(strs []string) string {
	// check if strs empty
	if len(strs) == 0 {
		return ""
	}
	for i := range strs[0] {
		c := strs[0][i]
		for j := range strs {
			if i == len(strs[j]) || strs[j][i] != c {
				return strs[0][0:i]
			}
		}
	}
	return strs[0]
}

//Approach 3: Divide and conquer
func longestCommonPrefix3(strs []string) string {
	if len(strs) == 0 {
		return ""
	}
	return dfs(strs, 0, len(strs)-1)
}

func dfs(strs []string, l int, r int) string {
	if l == r {
		return strs[l]
	}
	mid := (l + r) / 2
	lcpLeft := dfs(strs, l, mid)
	lcpRight := dfs(strs, mid+1, r)
	return commonPrefix(lcpLeft, lcpRight)
}

func commonPrefix(left string, right string) string {
	var min int
	if len(left) > len(right) {
		min = len(right)
	} else {
		min = len(left)
	}
	for i := 0; i < min; i++ {
		if left[i] != right[i] {
			return left[0:i]
		}
	}
	return left[0:min]
}

//Approach 4: Binary search
func longestCommonPrefix4(strs []string) string {
	if len(strs) == 0 {
		return ""
	}
	minLen := 1<<32 - 1
	for _, v := range strs {
		if minLen > len(v) {
			minLen = len(v)
		}
	}
	low := 1
	high := minLen

	for low <= high {
		middle := (low + high) / 2
		if isCommonPrefix(strs, middle) {
			low = middle + 1
		} else {
			low = middle - 1
		}
	}
	return strs[0][0 : (low+high)/2]
}

func isCommonPrefix(strs []string, len int) bool {
	str1 := strs[0][0:len]
	for _, v := range strs {
		if !strings.HasPrefix(v, str1) {
			return false
		}
	}
	return true
}

func main() {
	strs := []string{"flower", "flow", "flo"}
	fmt.Println(longestCommonPrefix2(strs))
	fmt.Println(longestCommonPrefix3(strs))
	fmt.Println(longestCommonPrefix4(strs))
}
