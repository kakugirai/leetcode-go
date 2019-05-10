package main

import "fmt"

func firstUniqChar(s string) int {
	alphabets := make(map[byte]int)
	for i := 0; i < len(s); i++ {
		if _, ok := alphabets[s[i]]; !ok {
			alphabets[s[i]] = i
		} else {
			alphabets[s[i]] = -1
		}
	}
	result := int(^uint(0) >> 1) // max int value
	for _, v := range alphabets {
		if v < result && v != -1 {
			result = v
		}
	}
	if result == int(^uint(0)>>1) {
		return -1
	}
	return result
}

func main() {
	fmt.Println(firstUniqChar("loveleetcode"))
}
