package main

import (
	"fmt"
	"math"
)

func findTheDifference(s string, t string) byte {
	sSum := 0
	tSum := 0
	for i := range s {
		sSum += int(s[i])
	}
	for i := range t {
		tSum += int(t[i])
	}
	return byte(math.Abs(float64(sSum - tSum)))
}

func main() {
	fmt.Println(findTheDifference("abcd", "abcde"))
}
