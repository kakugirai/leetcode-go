package main

import (
	"fmt"
	"unicode"
)

func licenseKeyFormatting(S string, K int) string {
	count := 0
	var result []rune
	i := len(S) - 1

	for i >= 0 {
		current := unicode.ToUpper(rune(S[i]))
		if current == '-' {
			i--
		} else if count != 0 && count%K == 0 {
			result = append([]rune{'-'}, result...)
			count = 0
		} else {
			result = append([]rune{current}, result...)
			count++
			i--
		}
	}
	return string(result)
}

func main() {
	S := "5F3Z-2e-9-w"
	K := 4
	fmt.Println(licenseKeyFormatting(S, K))
}
