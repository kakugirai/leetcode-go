package main

import (
	"fmt"
	"sort"
)

func groupAnagrams(strs []string) [][]string {
	var result [][]string
	m := make(map[string][]string)
	for _, current := range strs {
		var c []byte
		for i := range current {
			c = append(c, current[i])
		}
		sort.Slice(c, func(i int, j int) bool { return c[i] < c[j] })
		sorted := string(c)
		if _, ok := m[sorted]; !ok {
			m[sorted] = []string{}
		}
		m[sorted] = append(m[sorted], current)
	}

	for _, v := range m {
		result = append(result, v)
	}
	return result
}

func main() {
	strs := []string{
		"eat", "tea", "tan", "ate", "nat", "bat",
	}
	fmt.Printf("%#v", groupAnagrams(strs))
}
