package main

import "fmt"

func findRepeatedDnaSequences(s string) []string {
	m := make(map[string]int)
	var result []string
	for i := 0; i <= len(s)-10; i++ {
		ss := s[i : i+10]
		m[ss]++
		if i, ok := m[ss]; ok && i == 2 {
			result = append(result, ss)
		}
	}
	return result
}

func main() {
	s := "AAAAAAAAAAA"
	fmt.Println(findRepeatedDnaSequences(s))
}
