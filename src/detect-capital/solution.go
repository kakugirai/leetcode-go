package main

import "fmt"

func detectCapitalUse(word string) bool {
	countUpper := 0
	for i := 0; i < len(word); i++ {
		if word[i] >= 'A' && word[i] <= 'Z' {
			countUpper++
		}
	}
	return countUpper == len(word) || countUpper == 0 || (countUpper == 1 && word[0] >= 'A' && word[0] <= 'Z')
}

func main() {
	word := "Google"
	fmt.Println(detectCapitalUse(word))
}
