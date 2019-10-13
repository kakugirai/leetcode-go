package main

import "fmt"

func isIsomorphic(s string, t string) bool {
	if len(s) != len(t) {
		return false
	}
	return checkIsomorphic(s, t) && checkIsomorphic(t, s)
}

func checkIsomorphic(s string, t string) bool {
	m := make(map[byte]byte)
	for i := range s {
		if _, ok := m[s[i]]; ok {
			if m[s[i]] != t[i] {
				return false
			}
		} else {
			m[s[i]] = t[i]
		}
	}
	return true
}

func main() {
	s := "ab"
	t := "aa"
	fmt.Println(isIsomorphic(s, t))
}
