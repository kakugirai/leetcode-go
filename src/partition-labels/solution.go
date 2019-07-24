package main

import "fmt"

func partitionLabels(S string) []int {
	m := make(map[byte][]int)
	for i := range S {
		if _, ok := m[S[i]]; !ok {
			m[S[i]] = []int{i, 0}
		}
	}
	for i := range S {
		if i > m[S[i]][1] {
			m[S[i]][1] = i
		}
	}

	var result []int
	current := m[S[0]]
	for i := 1; i < len(S); i++ {
		if m[S[i]][0] > current[1] {
			result = append(result, current[1]-current[0]+1)
			current = m[S[i]]
		}
		if m[S[i]][0] > current[0] && m[S[i]][1] > current[1] {
			current[1] = m[S[i]][1]
		}
	}
	result = append(result, current[1]-current[0]+1)
	return result
}

func main() {
	S := "ababcbacadefegdehijhklij"
	fmt.Printf("%#v", partitionLabels(S))
}
