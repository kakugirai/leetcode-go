package main

import (
	"fmt"
	"strings"
)

func multiply(num1 string, num2 string) string {
	m, n := len(num1), len(num2)
	pos := make([]int, m+n)

	for i := m - 1; i >= 0; i-- {
		for j := n - 1; j >= 0; j-- {
			mul := int((num1[i] - '0') * (num2[j] - '0'))
			p1, p2 := i+j, i+j+1
			sum := mul + pos[p2]
			pos[p1] += sum / 10
			pos[p2] = sum % 10
		}
	}

	var sb strings.Builder
	for _, v := range pos {
		if !(sb.Len() == 0 && v == 0) {
			sb.WriteByte(byte(v) + '0')
		}
	}
	if sb.Len() == 0 {
		return "0"
	}
	return sb.String()
}

func main() {
	fmt.Println(multiply("408", "5"))
}
