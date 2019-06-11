package main

import (
	"fmt"
)

func removeOuterParentheses(S string) string {
	m := map[byte]int{'(': 1, ')': -1}
	var sum int
	var arr []string
	var ptr int

	for i := 0; i < len(S)-1; i++ {
		sum += m[S[i]]
		if sum == 0 {
			arr = append(arr, S[ptr:i+1])
			ptr = i + 1
		}
	}
	arr = append(arr, S[ptr:])

	var str string
	for _, v := range arr {
		str = str + v[1:len(v)-1]
	}
	return str
}

func main() {
	S := "(()())(())(()(()))"
	fmt.Println(removeOuterParentheses(S))
}
