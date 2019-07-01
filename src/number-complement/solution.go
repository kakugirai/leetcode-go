package main

import (
	"fmt"
	"strconv"
)

func findComplement(num int) int {
	var result string
	for num > 0 {
		result = strconv.Itoa(num%2^1) + result
		num /= 2
	}
	a, err := strconv.ParseInt(result, 2, 64)
	if err == nil {
		return int(a)
	}
	return -1
}

func main() {
	fmt.Println(findComplement(5))
}
