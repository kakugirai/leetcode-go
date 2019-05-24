package main

import "fmt"

func isAnagram(s string, t string) bool {
	if len(s) != len(t) {
		return false
	}
	var num [26]int
	for i := 0; i < len(s); i++ {
		num[s[i]-'a']++
		num[t[i]-'a']--
	}
	for _, v := range num {
		if v != 0 {
			return false
		}
	}
	return true
}

func main() {
	s := "abcdcc"
	t := "dcbacc"
	fmt.Println(isAnagram(s, t))
}
