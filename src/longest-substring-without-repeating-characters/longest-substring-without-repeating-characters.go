package main

import (
	"fmt"
)

func lengthOfLongestSubstring(s string) int {
	var ss []string

	m := make(map[byte]struct{})
	var sb []byte
	for i := range s {
		if _, ok := m[s[i]]; !ok {
			m[s[i]] = struct{}{}
			sb = append(sb, s[i])
		} else {
			ss = append(ss, string(sb))
			for j := range sb {
				if sb[j] == s[i] {
					for k := range sb[:j] {
						delete(m, sb[k])
					}
					sb = sb[j+1:]
					break
				}
			}
			sb = append(sb, s[i])
		}
	}
	ss = append(ss, string(sb))

	var current string
	for i := range ss {
		if len(ss[i]) > len(current) {
			current = ss[i]
		}
	}

	return len(current)
}

func main() {
	s := "asjrgapa"
	fmt.Println(lengthOfLongestSubstring(s))
}
