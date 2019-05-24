package main

import "fmt"

func numJewelsInStones(J string, S string) int {
	m := make(map[byte]bool)
	for i := range J {
		m[J[i]] = true
	}
	jewels := 0
	for i := range S {
		if _, ok := m[S[i]]; ok {
			jewels++
		}
	}
	return jewels
}

func main() {
	J := "aA"
	S := "aAAbbbb"
	fmt.Println(numJewelsInStones(J, S))
}
