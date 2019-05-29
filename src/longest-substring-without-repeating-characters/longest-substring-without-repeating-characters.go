package main

import (
	"bytes"
	"fmt"
)

func lengthOfLongestSubstring(s string) int {
	var result int
	var sb []byte
	m := make(map[byte]struct{})
	for i := range s {
		if _, ok := m[s[i]]; !ok {
			m[s[i]] = struct{}{}
			sb = append(sb, s[i])
		} else {
			if len(sb) > result {
				result = len(sb)
			}
			idx := bytes.IndexByte(sb, s[i])
			for k := 0; k < idx; k++ {
				delete(m, sb[k])
			}
			sb = sb[idx+1:]
			sb = append(sb, s[i])
		}
	}
	if len(sb) > result {
		result = len(sb)
	}
	return result
}

func main() {
	s := "asjrgapa"
	fmt.Println(lengthOfLongestSubstring(s))
}
