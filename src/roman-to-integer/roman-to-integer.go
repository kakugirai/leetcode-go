package main

import "fmt"

func main() {
	fmt.Println(romanToInt("MCMXCIV"))
}

func romanToInt(s string) int {
	m := map[string]int{"I": 1, "V": 5, "X": 10, "L": 50, "C": 100, "D": 500, "M": 1000}
	ans := 0
	for i, v := range s[:len(s)-1] {
		c := string(v)
		cNext := string(s[i+1])
		if m[c] < m[cNext] {
			ans -= m[c]
		} else {
			ans += m[c]
		}
	}
	ans += m[string(s[len(s)-1:])]
	return ans
}
