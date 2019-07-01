package main

import (
	"fmt"
	"strings"
)

func intToRoman(num int) string {
	values := []int{1000, 900, 500, 400, 100, 90, 50, 40, 10, 9, 5, 4, 1}
	arr := []string{"M", "CM", "D", "CD", "C", "XC", "L", "XL", "X", "IX", "V", "IV", "I"}
	sb := strings.Builder{}
	for i := 0; i < len(values); i++ {
		for num >= values[i] {
			num -= values[i]
			sb.WriteString(arr[i])
		}
	}
	return sb.String()
}

func main() {
	fmt.Println(intToRoman(1994))
}
