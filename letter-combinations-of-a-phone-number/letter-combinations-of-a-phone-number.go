package main

import (
	"fmt"
)

func letterCombinations(digits string) []string {
	var result []string
	if len(digits) == 0 {
		return result
	}

	m := []string{
		"0",
		"1",
		"abc",
		"def",
		"ghi",
		"jkl",
		"mno",
		"pqrs",
		"tuv",
		"wxyz",
	}
	letterCombinationsRecursive(&result, digits, "", 0, m)
	return result
}

func letterCombinationsRecursive(result *[]string, digits string, current string, index int, m []string) {
	if index == len(digits) {
		*result = append(*result, current)
		return
	}

	letters := m[int(digits[index]-'0')]
	for i := 0; i < len(letters); i++ {
		letterCombinationsRecursive(result, digits, current+string(letters[i]), index+1, m)
	}
}

func main() {
	fmt.Println(letterCombinations("23"))
}
