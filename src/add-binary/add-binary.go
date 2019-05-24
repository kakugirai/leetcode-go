package main

import (
	"fmt"
	"strconv"
)

func addBinary(a string, b string) string {
	i := len(a) - 1
	j := len(b) - 1
	carry := 0
	result := ""
	for i >= 0 || j >= 0 {
		sum := carry
		if i >= 0 {
			sum += int(a[i]) - '0'
			i--
		}
		if j >= 0 {
			sum += int(b[j]) - '0'
			j--
		}
		result = strconv.Itoa(sum%2) + result
		carry = sum / 2
	}
	if carry > 0 {
		result = "1" + result
	}
	return result
}

func main() {
	fmt.Println(addBinary("10110", "1101"))
}
