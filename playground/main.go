package main

import (
	"fmt"
)

func printReverse(str []byte) {
	helper(0, str)
}

func helper(idx int, str []byte) {
	if idx >= len(str) {
		return
	}
	helper(idx+1, str)
	fmt.Println(str[idx])
}

func main() {
	fmt.Println(int(1<<63 - 1))
	printReverse([]byte{1, 2, 3, 4})
}
