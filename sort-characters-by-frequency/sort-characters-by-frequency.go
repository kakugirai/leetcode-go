package main

import (
	"fmt"
	"sort"
)

type segments []string

func (s segments) Len() int {
	return len(s)
}

func (s segments) Less(i, j int) bool {
	if len(s[i]) == len(s[j]) {
		return s[i] < s[j]
	}
	return len(s[i]) > len(s[j])
}

func (s segments) Swap(i, j int) {
	s[i], s[j] = s[j], s[i]
}

func frequencySort(s string) string {
	freq := make([]int, 'z'+1)
	for i := range s {
		freq[s[i]]++
	}

	// strs := [ttt ee r]
	var strs []string
	for i := range freq {
		if freq[i] == 0 {
			continue
		}

		bytes := make([]byte, freq[i])
		for j := range bytes {
			bytes[j] = byte(i)
		}
		strs = append(strs, string(bytes))
	}

	// sort.Interface is implemented for segments
	sort.Sort(segments(strs))

	res := ""
	for _, s := range strs {
		res += s
	}

	return res
}

func main() {
	fmt.Println(frequencySort("tttree"))
}
