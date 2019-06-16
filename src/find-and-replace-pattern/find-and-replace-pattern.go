package main

import (
	"fmt"
)

func findAndReplacePattern(words []string, pattern string) []string {
	var result []string
	for _, word := range words {
		if match(word, pattern) {
			result = append(result, word)
		}
	}
	return result
}

func match(word string, pattern string) bool {
	m1 := make(map[byte]byte)
	m2 := make(map[byte]byte)

	for i := range word {
		w := word[i]
		p := pattern[i]
		if _, ok := m1[w]; !ok {
			m1[w] = p
		}
		if _, ok := m2[p]; !ok {
			m2[p] = w
		}
		if m1[w] != p || m2[p] != w {
			return false
		}
	}
	return true
}

func main() {
	words := []string{"abc", "deq", "mee", "aqq", "dkd", "ccc"}
	pattern := "abb"
	fmt.Println(findAndReplacePattern(words, pattern))
}
