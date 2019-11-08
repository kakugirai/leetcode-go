package main

import "fmt"

func minimumSwap(s1 string, s2 string) int {
	if len(s1) == 0 {
		return 0
	}
	m := make(map[int]int)
	for i := range s1 {
		if s1[i] == 'x' && s2[i] == 'y' {
			m[0] += 1
		}
		if s1[i] == 'y' && s2[i] == 'x' {
			m[1] += 1
		}
	}
	if (m[0]-m[1])%2 != 0 {
		return -1
	}
	return m[0]%2 + m[0]/2 + m[1]%2 + m[1]/2
}

func main() {
	s1 := "xxyyxyxyxx"
	s2 := "xyyxyxxxyx"
	fmt.Println(minimumSwap(s1, s2))
}
